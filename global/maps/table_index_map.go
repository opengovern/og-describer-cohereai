package maps

import (
	"github.com/opengovern/og-describer-cohereai/discovery/pkg/es"
)

var ResourceTypesToTables = map[string]string{
  "CohereAI/Connectors": "cohereai_connectors",
  "CohereAI/Models": "cohereai_models",
  "CohereAI/Datasets": "cohereai_datasets",
  "CohereAI/FineTunedModel": "cohereai_fine_tuned_models",
  "CohereAI/EmbedJob": "cohereai_embed_jobs",
}

var ResourceTypeToDescription = map[string]interface{}{
  "CohereAI/Connectors": opengovernance.Connector{},
  "CohereAI/Models": opengovernance.Model{},
  "CohereAI/Datasets": opengovernance.Dataset{},
  "CohereAI/FineTunedModel": opengovernance.FineTunedModel{},
  "CohereAI/EmbedJob": opengovernance.EmbedJob{},
}

var TablesToResourceTypes = map[string]string{
  "cohereai_connectors": "CohereAI/Connectors",
  "cohereai_models": "CohereAI/Models",
  "cohereai_datasets": "CohereAI/Datasets",
  "cohereai_fine_tuned_models": "CohereAI/FineTunedModel",
  "cohereai_embed_jobs": "CohereAI/EmbedJob",
}
