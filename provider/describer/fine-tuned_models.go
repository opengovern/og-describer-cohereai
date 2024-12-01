package describer
import (
	"context"
	"encoding/json"
	// "fmt"

	// "fmt"
	"net/http"
	"net/url"
	"sync"

	"github.com/opengovern/og-describer-cohereai/pkg/sdk/models"
	"github.com/opengovern/og-describer-cohereai/provider/model"
)


func ListFineTunedModels(ctx context.Context, handler *CohereAIAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
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

func processFineTunedModels(ctx context.Context, handler *CohereAIAPIHandler, cohereAiChan chan<- models.Resource, wg *sync.WaitGroup) {
	var fineTunedModelResponse model.ListFineTunedModelsResponse
	var fineTunedModels []model.FineTunedModelDescription
	var resp *http.Response
	baseURL := "https://api.cohere.com/v1/finetuning/finetuned-models"
	
	requestFunc := func(req *http.Request) (*http.Response, error) {
		var e error
		pageToken := ""
		for {
		params := url.Values{}
		params.Set("page_size", "1000")
		if(pageToken != ""){
			params.Set("pageToken", pageToken)
		}

		
		finalURL := baseURL + "?" + params.Encode()
		req, err := http.NewRequest("GET", finalURL, nil)
		if err != nil {
			return  nil,e
		}
		
		resp, e = handler.Client.Do(req)
		// fmt.Printf(json.NewDecoder(resp.Body))
		if e = json.NewDecoder(resp.Body).Decode(&fineTunedModelResponse); e != nil {
			return nil, e
		}
		fineTunedModels = append(fineTunedModels, fineTunedModelResponse.FinetunedModels...)
		
		if(fineTunedModelResponse.NextPageToken == ""){
			break
		}
		pageToken = fineTunedModelResponse.NextPageToken

	}
	return resp, e
	}

	err := handler.DoRequest(ctx,  &http.Request{}, requestFunc)
	if err != nil {
		return
	}
	for _, mod := range fineTunedModels {
		wg.Add(1)
		go func(mod model.FineTunedModelDescription) {
			defer wg.Done()
			value := models.Resource{
				ID:   mod.ID,
				Name: mod.Name,
				Description: JSONAllFieldsMarshaller{
					Value: mod,
				},
			}
			cohereAiChan <- value
		}(mod)
	}
}



