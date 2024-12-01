package steampipe

import (
	"github.com/opengovern/og-describer-cohereai/pkg/sdk/es"
)

var Map = map[string]string{
  "CohereAI/Classify": "cohereai_classify",
  "CohereAI/DetectLanguage": "cohereai_detect_language",
  "CohereAI/Detokenize": "cohereai_detokenize",
  "CohereAI/Embed": "cohereai_embed",
  "CohereAI/Generation": "cohereai_generation",
  "CohereAI/Summarize": "cohereai_summarize",
  "CohereAI/Tokenize": "cohereai_tokenize",
}

var DescriptionMap = map[string]interface{}{
  "CohereAI/Classify": opengovernance.Classify{},
  "CohereAI/DetectLanguage": opengovernance.DetectLanguage{},
  "CohereAI/Detokenize": opengovernance.Detokenize{},
  "CohereAI/Embed": opengovernance.Embed{},
  "CohereAI/Generation": opengovernance.Generation{},
  "CohereAI/Summarize": opengovernance.Summarize{},
  "CohereAI/Tokenize": opengovernance.Tokenize{},
}

var ReverseMap = map[string]string{
  "cohereai_classify": "CohereAI/Classify",
  "cohereai_detect_language": "CohereAI/DetectLanguage",
  "cohereai_detokenize": "CohereAI/Detokenize",
  "cohereai_embed": "CohereAI/Embed",
  "cohereai_generation": "CohereAI/Generation",
  "cohereai_summarize": "CohereAI/Summarize",
  "cohereai_tokenize": "CohereAI/Tokenize",
}
