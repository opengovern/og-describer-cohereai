package describers

import (
	"context"
	"encoding/json"
	"fmt"

	// "fmt"

	// "fmt"
	"net/http"
	"net/url"
	"sync"

	"github.com/opengovern/og-describer-cohereai/discovery/pkg/models"
	model "github.com/opengovern/og-describer-cohereai/discovery/provider"
)

func ListFineTunedModels(ctx context.Context, handler *model.CohereAIAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	cohereaiChan := make(chan models.Resource)

	go func() {
		processFineTunedModels(ctx, handler, cohereaiChan, &wg)
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

// page_size and page token

func processFineTunedModels(ctx context.Context, handler *model.CohereAIAPIHandler, cohereAiChan chan<- models.Resource, wg *sync.WaitGroup) {
	var fineTunedModelResponse model.ListFineTunedModelsResponse
	var fineTunedModels []model.FineTunedModelDescription
	var resp *http.Response
	baseURL := "https://api.cohere.com/v1/finetuning/finetuned-models"
	req, err1 := http.NewRequest("GET", baseURL, nil)
	if err1 != nil {
		return
	}
	requestFunc := func(req *http.Request) (*http.Response, error) {
		var e error
		pageToken := ""
		for {
			params := url.Values{}
			params.Set("page_size", "1000")
			if pageToken != "" {
				params.Set("pageToken", pageToken)
			}

			finalURL := baseURL + "?" + params.Encode()
			req, err := http.NewRequest("GET", finalURL, nil)
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", handler.APIKey))
			if handler.ClientName != "" {
				req.Header.Set("X-Client-Name", handler.ClientName)
			}
			if err != nil {
				return nil, e
			}

			resp, e = handler.Client.Do(req)
			// fmt.Printf(json.NewDecoder(resp.Body))
			if e = json.NewDecoder(resp.Body).Decode(&fineTunedModelResponse); e != nil {
				return nil, e
			}
			fineTunedModels = append(fineTunedModels, fineTunedModelResponse.FinetunedModels...)

			if fineTunedModelResponse.NextPageToken == "" {
				break
			}
			pageToken = fineTunedModelResponse.NextPageToken

		}
		return resp, e
	}

	err := handler.DoRequest(ctx, req, requestFunc)
	if err != nil {
		return
	}
	for _, mod := range fineTunedModels {
		wg.Add(1)
		go func(mod model.FineTunedModelDescription) {
			defer wg.Done()
			value := models.Resource{
				ID:          mod.ID,
				Name:        mod.Name,
				Description: mod,
			}
			cohereAiChan <- value
		}(mod)
	}
}

func GetFineTunedModel(ctx context.Context, handler *model.CohereAIAPIHandler, fineTunedModelID string) (*models.Resource, error) {
	var fineTunedModelResponse model.FineTunedModelDescription
	var fineTunedModel model.FineTunedModelDescription
	baseURL := "https://api.cohere.com/v1/finetuning/finetuned-models"
	finalURL := baseURL + "/" + fineTunedModelID
	req, err := http.NewRequest("GET", finalURL, nil)
	if err != nil {
		return nil, err
	}
	requestFunc := func(req *http.Request) (*http.Response, error) {
		var e error
		resp, e := handler.Client.Do(req)
		if e = json.NewDecoder(resp.Body).Decode(&fineTunedModelResponse); e != nil {
			return nil, e
		}
		fineTunedModel = fineTunedModelResponse
		return resp, e
	}
	err = handler.DoRequest(ctx, req, requestFunc)
	if err != nil {
		return nil, err
	}
	value := models.Resource{
		ID:          fineTunedModel.ID,
		Name:        fineTunedModel.Name,
		Description: fineTunedModel,
	}
	return &value, nil
}
