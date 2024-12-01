package steampipe

import (
	"github.com/opengovern/og-describer-cohereai/pkg/sdk/es"
)

var Map = map[string]string{
  "CohereAI/Connectors": "cohereai_connectors",
  "CohereAI/Model": "cohereai_models",
  "CohereAI/Datasets": "cohereai_datasets",
}

var DescriptionMap = map[string]interface{}{
  "CohereAI/Connectors": opengovernance.Connector{},
  "CohereAI/Model": opengovernance.Model{},
  "CohereAI/Datasets": opengovernance.Dataset{},
}

var ReverseMap = map[string]string{
  "cohereai_connectors": "CohereAI/Connectors",
  "cohereai_models": "CohereAI/Model",
  "cohereai_datasets": "CohereAI/Datasets",
}
