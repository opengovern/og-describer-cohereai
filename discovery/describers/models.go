package describers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"sync"

	"github.com/opengovern/og-describer-cohereai/discovery/pkg/models"
	model "github.com/opengovern/og-describer-cohereai/discovery/provider"
)

func ListModels(ctx context.Context, handler *model.CohereAIAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	cohereaiChan := make(chan models.Resource)
	

	go func() {
		processModels(ctx, handler, cohereaiChan, &wg)
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

func processModels(ctx context.Context, handler *model.CohereAIAPIHandler, cohereAiChan chan<- models.Resource, wg *sync.WaitGroup) {
	var modelResponse model.ModelListResponse
	var models_arr []model.ModelDescription
	var resp *http.Response
	baseURL := "https://api.cohere.com/v1/models"
	req,err1 := http.NewRequest("GET", baseURL, nil)
	if(err1 != nil){
		return 
	}
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
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", handler.APIKey))
		if handler.ClientName != "" {
			req.Header.Set("X-Client-Name", handler.ClientName)
		}
		if err != nil {
			return  nil,e
		}
		
		resp, e = handler.Client.Do(req)
		// fmt.Printf(json.NewDecoder(resp.Body))
		if e = json.NewDecoder(resp.Body).Decode(&modelResponse); e != nil {
			return nil, e
		}
		models_arr = append(models_arr, modelResponse.Models...)
		
		if(modelResponse.NextPageToken == ""){
			break
		}
		pageToken = modelResponse.NextPageToken

	}
	return resp, e
	}

	err := handler.DoRequest(ctx,  req, requestFunc)
	if err != nil {
		return
	}
	for _, mod := range models_arr {
		wg.Add(1)
		go func(mod model.ModelDescription) {
			defer wg.Done()
			value := models.Resource{
				ID:   mod.Name,
				Name: mod.Name,
				Description: JSONAllFieldsMarshaller{
					Value: mod,
				},
			}
			cohereAiChan <- value
		}(mod)
	}
}

func GetModel(ctx context.Context, handler *model.CohereAIAPIHandler, modelName string) (*models.Resource, error) {
	var modelResponse model.ModelDescription
	baseURL := "https://api.cohere.com/v1/models"
	
	
	finalURL := baseURL + "/" + modelName
	req, err := http.NewRequest("GET", finalURL, nil)
	if err != nil {
		return &models.Resource{}, err
	}
	requestFunc := func(req *http.Request) (*http.Response, error) {
		var e error
		
		resp, e := handler.Client.Do(req)
		if e = json.NewDecoder(resp.Body).Decode(&modelResponse); e != nil {
			return nil, e
		}
		
		return resp, e
	}
	err = handler.DoRequest(ctx, req, requestFunc)
	if err != nil {
		return &models.Resource{}, err
	}
	value := models.Resource{
		ID:   modelResponse.Name,
		Name: modelResponse.Name,
		Description: JSONAllFieldsMarshaller{
			Value: modelResponse,
		},
	}
	return &value, nil
}