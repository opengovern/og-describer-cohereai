package cohere

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-cohereai/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableCohereModels(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "cohereai_models",
		Description: "Cohere Ai list of models.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListModel,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			Hydrate:    opengovernance.GetModel,
		},
		Columns: commonColumns([]*plugin.Column{
			// Top columns
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Name"),
				Description: "Name of the model."},
			{
				Name:        "endpoints",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Endpoints"),
				Description: "Endpoints of the model."},
			{
				Name:        "tokenizer_url",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.TokenizerUrl"),
				Description: "Tokenizer URL of the model."},
			{
				Name:        "default_endpoints",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DefaultEndpoints"),
				Description: "Default endpoints of the model."},
			
	
			
			
		}),
	}
}
