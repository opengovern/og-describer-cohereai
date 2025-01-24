package main

import (
	"encoding/json"
	"github.com/opengovern/og-describer-cohereai/platform/constants"
	"github.com/jackc/pgtype"
	"github.com/opengovern/og-describer-cohereai/global"
	"github.com/opengovern/og-describer-cohereai/global/maps"
	"github.com/opengovern/og-util/pkg/integration"
	"github.com/opengovern/og-util/pkg/integration/interfaces"
)

type Integration struct{}

func (i *Integration) GetConfiguration() (interfaces.IntegrationConfiguration, error) {
	return interfaces.IntegrationConfiguration{
		NatsScheduledJobsTopic:   global.JobQueueTopic,
		NatsManualJobsTopic:      global.JobQueueTopicManuals,
		NatsStreamName:           global.StreamName,
		NatsConsumerGroup:        global.ConsumerGroup,
		NatsConsumerGroupManuals: global.ConsumerGroupManuals,

		SteampipePluginName: "cohereai",

		UISpec:   constants.UISpec,
		Manifest: constants.Manifest,
		SetupMD:  constants.SetupMd,

		DescriberDeploymentName: constants.DescriberDeploymentName,
		DescriberRunCommand:     constants.DescriberRunCommand,
	}, nil
}

func (i *Integration) HealthCheck(jsonData []byte, providerId string, labels map[string]string, annotations map[string]string) (bool, error) {
	var credentials global.IntegrationCredentials
	err := json.Unmarshal(jsonData, &credentials)
	if err != nil {
		return false, err
	}

	isHealthy, err := CohereAIIntegrationHealthcheck(credentials.APIKey)
	return isHealthy, err
}

func (i *Integration) DiscoverIntegrations(jsonData []byte) ([]integration.Integration, error) {
	var credentials global.IntegrationCredentials
	err := json.Unmarshal(jsonData, &credentials)
	if err != nil {
		return nil, err
	}
	var integrations []integration.Integration
	connectors, err1 := CohereAIIntegrationDiscovery(credentials.APIKey)
	if err1 != nil {
		return nil, err1
	}
	labels := map[string]string{
		"ClientName": credentials.ClientName,
	}
	if len(connectors) > 0 {
		labels["OrganizationID"] = connectors[0].OrganizationID
	}
	labelsJsonData, err := json.Marshal(labels)
	if err != nil {
		return nil, err
	}
	integrationLabelsJsonb := pgtype.JSONB{}
	err = integrationLabelsJsonb.Set(labelsJsonData)
	if err != nil {
		return nil, err
	}
	// for in esponse
	for _, connector := range connectors {
		integrations = append(integrations, integration.Integration{
			ProviderID: connector.ID,
			Name:       connector.Name,
			Labels:     integrationLabelsJsonb,
		})
	}

	return integrations, nil
}

func (i *Integration) GetResourceTypesByLabels(labels map[string]string) ([]interfaces.ResourceTypeConfiguration, error) {
	var resourceTypesMap  []interfaces.ResourceTypeConfiguration
	for _, resourceType := range maps.ResourceTypesList {
		var resource  interfaces.ResourceTypeConfiguration
		if v, ok := maps.ResourceTypeConfigs[resourceType]; ok {
			resource.Description =v.Description
			resource.Params =v.Params
			resource.Name = v.Name
			resource.IntegrationType = v.IntegrationType
			resource.Table =  maps.ResourceTypesToTables[v.Name]
			resourceTypesMap = append(resourceTypesMap, resource)
			
		} 
	}
	return resourceTypesMap, nil
}

func (i *Integration) GetResourceTypeFromTableName(tableName string) (string, error) {
	if v, ok := maps.TablesToResourceTypes[tableName]; ok {
		return v, nil
	}

	return "", nil
}

func (i *Integration) GetIntegrationType() (integration.Type, error) {
	return constants.IntegrationName, nil
}

func (i *Integration) ListAllTables() (map[string][]interfaces.CloudQLColumn, error) {
	plugin := global.Plugin()
	tables := make(map[string][]interfaces.CloudQLColumn)
	for tableKey, table := range plugin.TableMap {
		columns := make([]interfaces.CloudQLColumn, 0, len(table.Columns))
		for _, column := range table.Columns {
			columns = append(columns, interfaces.CloudQLColumn{Name: column.Name, Type: column.Type.String()})
		}
		tables[tableKey] = columns
	}

	return tables, nil
}

func (i *Integration) Ping() error {
	return nil
}
