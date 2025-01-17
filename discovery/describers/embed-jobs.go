package describers

import (
	"context"
	"encoding/json"
	// "fmt"

	// "fmt"
	"net/http"
	"sync"

	"github.com/opengovern/og-describer-cohereai/discovery/pkg/models"
	model "github.com/opengovern/og-describer-cohereai/discovery/provider"
)

func ListEmbedJobs(ctx context.Context, handler *CohereAIAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	cohereaiChan := make(chan models.Resource)

	go func() {
		processEmbedJobs(ctx, handler, cohereaiChan, &wg)
		wg.Wait()
		close(cohereaiChan)
	}()
	var values []models.Resource
	for value := range cohereaiChan {
		if stream != nil {
			if err := (*stream)(value); err != nil {
				return nil, err
			}
		} else {
			values = append(values, value)
		}
	}
	return values, nil
}

func processEmbedJobs(ctx context.Context, handler *CohereAIAPIHandler, cohereAiChan chan<- models.Resource, wg *sync.WaitGroup) {
	var embedJobResponse model.ListEmbedJobsResponse
	var embedJobs []model.EmbedJobDescription
	var resp *http.Response
	baseURL := "https://api.cohere.com/v1/embed-jobs"

	finalURL := baseURL
	req, err := http.NewRequest("GET", finalURL, nil)
	if err != nil {
		return
	}
	requestFunc := func(req *http.Request) (*http.Response, error) {
		var e error

		resp, e = handler.Client.Do(req)
		// fmt.Printf(json.NewDecoder(resp.Body))
		if e = json.NewDecoder(resp.Body).Decode(&embedJobResponse); e != nil {
			return nil, e
		}
		embedJobs = append(embedJobs, embedJobResponse.EmbedJobs...)

		return resp, nil
	}

	err = handler.DoRequest(ctx, req, requestFunc)
	if err != nil {
		return
	}
	for _, mod := range embedJobs {
		wg.Add(1)
		go func(mod model.EmbedJobDescription) {
			defer wg.Done()
			value := models.Resource{
				ID:          mod.JobID,
				Name:        mod.Name,
				Description: mod,
			}
			cohereAiChan <- value
		}(mod)
	}

}

func GetEmbedJob(ctx context.Context, handler *CohereAIAPIHandler, embedJobID string) (*models.Resource, error) {
	var embedJobResponse model.EmbedJobDescription
	var embedJob model.EmbedJobDescription
	baseURL := "https://api.cohere.com/v1/embed-jobs"
	finalURL := baseURL + "/" + embedJobID
	req, err := http.NewRequest("GET", finalURL, nil)
	if err != nil {
		return nil, err
	}
	requestFunc := func(req *http.Request) (*http.Response, error) {
		var e error
		resp, e := handler.Client.Do(req)
		if e = json.NewDecoder(resp.Body).Decode(&embedJobResponse); e != nil {
			return nil, e
		}
		embedJob = embedJobResponse
		return resp, e
	}
	err = handler.DoRequest(ctx, req, requestFunc)
	if err != nil {
		return nil, err
	}
	value := &models.Resource{
		ID:          embedJob.JobID,
		Name:        embedJob.Name,
		Description: embedJob,
	}
	return value, nil
}
