package provider

import (
	"errors"
	"time"

	model "github.com/opengovern/og-describer-cohereai/pkg/sdk/models"
	"github.com/opengovern/og-describer-cohereai/provider/configs"
	"github.com/opengovern/og-util/pkg/describe/enums"
	"github.com/opengovern/og-describer-cohereai/provider/describer"

	"golang.org/x/net/context"
	"golang.org/x/time/rate"
)

// DescribeListByCohereAI A wrapper to pass cohereAI authorization to describer functions
func DescribeListByCohereAI(describe func(context.Context, *describer.CohereAIAPIHandler, *model.StreamSender) ([]model.Resource, error)) model.ResourceDescriber {
	return func(ctx context.Context, cfg configs.IntegrationCredentials, triggerType enums.DescribeTriggerType, additionalParameters map[string]string, stream *model.StreamSender) ([]model.Resource, error) {
		ctx = describer.WithTriggerType(ctx, triggerType)

		var err error
		// Check for the token
		if cfg.APIKey == "" {
			return nil, errors.New("token must be configured")
		}

		cohereAPIHandler := describer.NewCohereAIAPIHandler(cfg.APIKey, rate.Every(time.Minute/200), 1, 10, 5, 5*time.Minute,cfg.ClientName)

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
func DescribeSingleByCohereAI(describe func(context.Context, *describer.CohereAIAPIHandler, string) (*model.Resource, error)) model.SingleResourceDescriber {
	return func(ctx context.Context, cfg configs.IntegrationCredentials, triggerType enums.DescribeTriggerType, additionalParameters map[string]string, resourceID string) (*model.Resource, error) {
		ctx = describer.WithTriggerType(ctx, triggerType)

		var err error
		// Check for the token
		if cfg.APIKey == "" {
			return nil, errors.New("token must be configured")
		}

		cohereAPIHandler := describer.NewCohereAIAPIHandler(cfg.APIKey, rate.Every(time.Minute/200), 1, 10, 5, 5*time.Minute,cfg.ClientName)

		// Get value from describer
		value, err := describe(ctx, cohereAPIHandler, resourceID)
		if err != nil {
			return nil, err
		}
		return value, nil
	}
}
