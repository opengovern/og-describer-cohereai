package provider
import (
	"github.com/opengovern/og-describer-cohereai/provider/describer"
	"github.com/opengovern/og-describer-cohereai/provider/configs"
	model "github.com/opengovern/og-describer-cohereai/pkg/sdk/models"
)
var ResourceTypes = map[string]model.ResourceType{

	"CohereAI/Connectors": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "CohereAI/Connectors",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeListByCohereAI(describer.ListConnectors),
		GetDescriber:         nil,
	},
}
