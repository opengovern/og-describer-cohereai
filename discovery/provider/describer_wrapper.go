package provider

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	model "github.com/opengovern/og-describer-cohereai/discovery/pkg/models"
	"github.com/opengovern/og-util/pkg/describe/enums"
	"golang.org/x/net/context"
	"golang.org/x/time/rate"
)

type CohereAIAPIHandler struct {
	Client       *http.Client
	APIKey        string
	RateLimiter  *rate.Limiter
	Semaphore    chan struct{}
	MaxRetries   int
	RetryBackoff time.Duration
	ClientName   string
}
func NewCohereAIAPIHandler(APIKey string, rateLimit rate.Limit, burst int, maxConcurrency int, maxRetries int, retryBackoff time.Duration,clientName string ) *CohereAIAPIHandler {
	return &CohereAIAPIHandler{
		Client:       http.DefaultClient,
		APIKey:        APIKey,
		RateLimiter:  rate.NewLimiter(rateLimit, burst),
		Semaphore:    make(chan struct{}, maxConcurrency),
		MaxRetries:   maxRetries,
		RetryBackoff: retryBackoff,
		ClientName: clientName,
	}
}
// DoRequest executes the openai API request with rate limiting, retries, and concurrency control.
func (h *CohereAIAPIHandler) DoRequest(ctx context.Context, req *http.Request, requestFunc func(req *http.Request) (*http.Response, error)) error {
	h.Semaphore <- struct{}{}
	// defer func() { <-h.Semaphore }()
	var resp *http.Response
	var err error
	for attempt := 0; attempt <= h.MaxRetries; attempt++ {
		// Wait based on rate limiter
		if err = h.RateLimiter.Wait(ctx); err != nil {
			return err
		}
		// Set request headers
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", h.APIKey))
		if h.ClientName != "" {
			req.Header.Set("X-Client-Name", h.ClientName)
		}
		// Execute the request function
		resp, err = requestFunc(req)
		if err == nil {
			return nil
		}
		// Handle rate limit errors
		if resp != nil && resp.StatusCode == http.StatusTooManyRequests {
			// Wait until reset time
			resetTime := resp.Header.Get("x-ratelimit-reset-requests")
			var resetDuration time.Duration
			if resetTime != "" {
				resetUnix, parseErr := strconv.ParseInt(resetTime, 10, 64)
				if parseErr == nil {
					resetDuration = time.Until(time.Unix(resetUnix, 0))
				}
			}
			// Set rate limiter value according to rate limit requests header
			var remainRequests int
			remainRequestsStr := resp.Header.Get("x-ratelimit-remaining-requests")
			if remainRequestsStr != "" {
				remainRequests, err = strconv.Atoi(remainRequestsStr)
				if err == nil {
					h.RateLimiter = rate.NewLimiter(rate.Every(resetDuration/time.Duration(remainRequests)), 1)
				}
			}
			if resetDuration > 0 {
				time.Sleep(resetDuration)
				continue
			}
			// Exponential backoff if headers are missing
			backoff := h.RetryBackoff * (1 << attempt)
			time.Sleep(backoff)
			continue
		}
		// Handle temporary network errors
		if isTemporary(err) {
			backoff := h.RetryBackoff * (1 << attempt)
			time.Sleep(backoff)
			continue
		}
		break
	}
	return err
}

// isTemporary checks if an error is temporary.
func isTemporary(err error) bool {
	if err == nil {
		return false
	}
	var netErr interface{ Temporary() bool }
	if errors.As(err, &netErr) {
		return netErr.Temporary()
	}
	return false
}

var (
	triggerTypeKey string = "trigger_type"
)
func WithTriggerType(ctx context.Context, tt enums.DescribeTriggerType) context.Context {
	return context.WithValue(ctx, triggerTypeKey, tt)
}

// DescribeListByCohereAI A wrapper to pass cohereAI authorization to describer functions
func DescribeListByCohereAI(describe func(context.Context, *CohereAIAPIHandler, *model.StreamSender) ([]model.Resource, error)) model.ResourceDescriber {
	return func(ctx context.Context, cfg model.IntegrationCredentials, triggerType enums.DescribeTriggerType, additionalParameters map[string]string, stream *model.StreamSender) ([]model.Resource, error) {
		ctx = WithTriggerType(ctx, triggerType)

		var err error
		// Check for the token
		if cfg.APIKey == "" {
			return nil, errors.New("token must be configured")
		}

		cohereAPIHandler := NewCohereAIAPIHandler(cfg.APIKey, rate.Every(time.Minute/200), 1, 10, 5, 5*time.Minute,cfg.ClientName)

		// Get values from describer
		var values []model.Resource
		result, err := describe(ctx, cohereAPIHandler, stream)
		if err != nil {
			return nil, err
		}
		values = append(values, result...)
		return values, nil
	}
}

// DescribeSingleByCohereAI A wrapper to pass cohereAI authorization to describer functions
func DescribeSingleByCohereAI(describe func(context.Context, *CohereAIAPIHandler, string) (*model.Resource, error)) model.SingleResourceDescriber {
	return func(ctx context.Context, cfg model.IntegrationCredentials, triggerType enums.DescribeTriggerType, additionalParameters map[string]string,resourceID string,stream *model.StreamSender) (*model.Resource, error) {
		ctx = WithTriggerType(ctx, triggerType)

		var err error
		// Check for the token
		if cfg.APIKey == "" {
			return nil, errors.New("API Key must be configured")
		}

		cohereAPIHandler := NewCohereAIAPIHandler(cfg.APIKey, rate.Every(time.Minute/200), 1, 10, 5, 5*time.Minute,cfg.ClientName)

		// Get value from describer
		value, err := describe(ctx, cohereAPIHandler, resourceID)
		if err != nil {
			return nil, err
		}
		return value, nil
	}
}
