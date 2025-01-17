package global

import "github.com/opengovern/og-util/pkg/integration"

const (
	IntegrationTypeLower = "cohereai"                    // example: aws, azure
	IntegrationName      = integration.Type("cohereai_project") // example: AWS_ACCOUNT, AZURE_SUBSCRIPTION
	OGPluginRepoURL      = "github.com/opengovern/og-describer-cohereai"                           // example: github.com/opengovern/og-describer-aws
)


type IntegrationCredentials struct {
	APIKey string `json:"api_key"`
	ClientName string `json:"client_name"`
}
