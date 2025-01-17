package describers

import (
	"context"
	"encoding/json"

	// "fmt"

	// "fmt"
	"net/http"
	"net/url"
	"sync"

	"github.com/opengovern/og-describer-cohereai/discovery/pkg/models"
	"github.com/opengovern/og-describer-cohereai/discovery/provider"
	model "github.com/opengovern/og-describer-cohereai/discovery/provider"
)

func ListConnectors(ctx context.Context, handler *provider.CohereAIAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	cohereaiChan := make(chan models.Resource)

	go func() {
		processConnectors(ctx, handler, cohereaiChan, &wg)
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

func processConnectors(ctx context.Context, handler *provider.CohereAIAPIHandler, cohereAiChan chan<- models.Resource, wg *sync.WaitGroup) {
	var connectorResponse model.ConnectorDescription
	var connectors []model.Connector
	var resp *http.Response
	baseURL := "https://api.cohere.com/v1/connectors"
	params := url.Values{}
	params.Set("limit", "100")

	finalURL := baseURL + "?" + params.Encode()
	req, err := http.NewRequest("GET", finalURL, nil)
	if err != nil {
		return
	}
	requestFunc := func(req *http.Request) (*http.Response, error) {
		var e error

		resp, e = handler.Client.Do(req)
		// fmt.Printf(json.NewDecoder(resp.Body))
		if e = json.NewDecoder(resp.Body).Decode(&connectorResponse); e != nil {
			return nil, e
		}
		connectors = append(connectors, connectorResponse.Connectors...)

		return resp, e
	}
	err = handler.DoRequest(ctx, req, requestFunc)
	if err != nil {
		return
	}
	for _, connector := range connectors {
		wg.Add(1)
		go func(connector model.Connector) {
			defer wg.Done()
			value := models.Resource{
				ID:          connector.ID,
				Name:        connector.Name,
				Description: connector,
			}
			cohereAiChan <- value
		}(connector)
	}
}

func GetConnector(ctx context.Context, handler *provider.CohereAIAPIHandler, connectorID string) (*models.Resource, error) {
	var connectorResponse model.ConnectorDetailResponse
	var connector model.Connector
	baseURL := "https://api.cohere.com/v1/connectors"
	finalURL := baseURL + "/" + connectorID
	req, err := http.NewRequest("GET", finalURL, nil)
	if err != nil {
		return nil, err
	}
	requestFunc := func(req *http.Request) (*http.Response, error) {
		var e error
		resp, e := handler.Client.Do(req)
		if e = json.NewDecoder(resp.Body).Decode(&connectorResponse); e != nil {
			return nil, e
		}
		connector = connectorResponse.Connector
		return resp, e
	}
	err = handler.DoRequest(ctx, req, requestFunc)
	if err != nil {
		return nil, err
	}
	value := &models.Resource{
		ID:          connector.ID,
		Name:        connector.Name,
		Description: connector,
	}
	return value, nil
}
