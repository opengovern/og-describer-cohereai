package provider
import (
	"github.com/opengovern/og-describer-cohereai/provider/describer"
	"github.com/opengovern/og-describer-cohereai/provider/configs"
	model "github.com/opengovern/og-describer-cohereai/pkg/sdk/models"
)
var ResourceTypes = map[string]model.ResourceType{

	"CohereAI/Classify": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "CohereAI/Classify",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeListByCohereAI(describer.ListClassifications),
		GetDescriber:         DescribeSingleByCohereAI(describer.GetClassification),
	},

	"CohereAI/DetectLanguage": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "CohereAI/DetectLanguage",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeListByCohereAI(describer.ListLanguages),
		GetDescriber:         DescribeSingleByCohereAI(describer.GetLanguageDetection),
	},

	"CohereAI/Detokenize": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "CohereAI/Detokenize",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeListByCohereAI(describer.ListDetokenizations),
		GetDescriber:         DescribeSingleByCohereAI(describer.GetDetokenization),
	},

	"CohereAI/Embed": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "CohereAI/Embed",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeListByCohereAI(describer.ListEmbeddings),
		GetDescriber:         DescribeSingleByCohereAI(describer.GetEmbedding),
	},

	"CohereAI/Generation": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "CohereAI/Generation",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeListByCohereAI(describer.ListGenerations),
		GetDescriber:         DescribeSingleByCohereAI(describer.GetGeneration),
	},

	"CohereAI/Summarize": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "CohereAI/Summarize",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeListByCohereAI(describer.ListSummarizations),
		GetDescriber:         DescribeSingleByCohereAI(describer.GetSummarization),
	},

	"CohereAI/Tokenize": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "CohereAI/Tokenize",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeListByCohereAI(describer.ListTokenizations),
		GetDescriber:         DescribeSingleByCohereAI(describer.GetTokenization),
	},
}
