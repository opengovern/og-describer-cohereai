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
		GetDescriber:         DescribeSingleByCohereAI(describer.GetConnector),
	},

	"CohereAI/Models": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "CohereAI/Models",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeListByCohereAI(describer.ListModels),
		GetDescriber:         DescribeSingleByCohereAI(describer.GetModel),
	},

	"CohereAI/Datasets": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "CohereAI/Datasets",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeListByCohereAI(describer.ListDatasets),
		GetDescriber:         DescribeSingleByCohereAI(describer.GetDataset),
	},

	"CohereAI/FineTunedModel": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "CohereAI/FineTunedModel",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeListByCohereAI(describer.ListFineTunedModels),
		GetDescriber:         DescribeSingleByCohereAI(describer.GetFineTunedModel),
	},

	"CohereAI/EmbedJob": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "CohereAI/EmbedJob",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeListByCohereAI(describer.ListEmbedJobs),
		GetDescriber:         DescribeSingleByCohereAI(describer.GetEmbedJob),
	},
}
