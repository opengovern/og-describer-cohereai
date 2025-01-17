package maps
import (
	"github.com/opengovern/og-describer-cohereai/discovery/describers"
	"github.com/opengovern/og-describer-cohereai/discovery/provider"
	"github.com/opengovern/og-describer-cohereai/platform/constants"
	"github.com/opengovern/og-util/pkg/integration/interfaces"
	model "github.com/opengovern/og-describer-cohereai/discovery/pkg/models"
)
var ResourceTypes = map[string]model.ResourceType{

	"CohereAI/Connectors": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "CohereAI/Connectors",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeListByCohereAI(describers.ListConnectors),
		GetDescriber:         provider.DescribeSingleByCohereAI(describers.GetConnector),
	},

	"CohereAI/Models": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "CohereAI/Models",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeListByCohereAI(describers.ListModels),
		GetDescriber:         provider.DescribeSingleByCohereAI(describers.GetModel),
	},

	"CohereAI/Datasets": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "CohereAI/Datasets",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeListByCohereAI(describers.ListDatasets),
		GetDescriber:         provider.DescribeSingleByCohereAI(describers.GetDataset),
	},

	"CohereAI/FineTunedModel": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "CohereAI/FineTunedModel",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeListByCohereAI(describers.ListFineTunedModels),
		GetDescriber:         provider.DescribeSingleByCohereAI(describers.GetFineTunedModel),
	},

	"CohereAI/EmbedJob": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "CohereAI/EmbedJob",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeListByCohereAI(describers.ListEmbedJobs),
		GetDescriber:         provider.DescribeSingleByCohereAI(describers.GetEmbedJob),
	},
}


var ResourceTypeConfigs = map[string]*interfaces.ResourceTypeConfiguration{

	"CohereAI/Connectors": {
		Name:         "CohereAI/Connectors",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"CohereAI/Models": {
		Name:         "CohereAI/Models",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"CohereAI/Datasets": {
		Name:         "CohereAI/Datasets",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"CohereAI/FineTunedModel": {
		Name:         "CohereAI/FineTunedModel",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"CohereAI/EmbedJob": {
		Name:         "CohereAI/EmbedJob",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},
}


var ResourceTypesList = []string{
  "CohereAI/Connectors",
  "CohereAI/Models",
  "CohereAI/Datasets",
  "CohereAI/FineTunedModel",
  "CohereAI/EmbedJob",
}