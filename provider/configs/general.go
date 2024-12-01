package configs

import "github.com/opengovern/og-util/pkg/integration"

const (
	IntegrationTypeLower = "cohereai"                    // example: aws, azure
	IntegrationName      = integration.Type("COHERE_AI") // example: AWS_ACCOUNT, AZURE_SUBSCRIPTION
	OGPluginRepoURL      = "github.com/opengovern/og-describer-cohere"                           // example: github.com/opengovern/og-describer-aws
)
