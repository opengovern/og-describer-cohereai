package cohere

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-cohereai/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableCohereFineTunedModels(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "cohereai_fine_tuned_models",
		Description: "Cohere Ai list of fine tuned models.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListFineTunedModel,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    opengovernance.GetFineTunedModel,
		},
		Columns: commonColumns([]*plugin.Column{
			// Top columns
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Name"),
				Description: "Name of the model."},
			
			{
				Name:        "default_endpoints",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DefaultEndpoints"),
				Description: "Default endpoints of the model."},
				{
					Name : "id",
					Type : proto.ColumnType_STRING,
					Transform: transform.FromField("Description.ID"),
					Description: "ID of the model.",
				},
				
				{
					Name : "created_at",
					Type : proto.ColumnType_TIMESTAMP,
					Transform: transform.FromField("Description.CreatedAt").Transform(transform.UnixToTimestamp),
					Description: "Timestamp of when the model was created.",
				},
				{
					Name : "updated_at",
					Type : proto.ColumnType_TIMESTAMP,
					Transform: transform.FromField("Description.UpdatedAt").Transform(transform.UnixToTimestamp),
					Description: "Timestamp of when the model was updated.",
				},
				{
					Name : "completed_at",
					Type : proto.ColumnType_TIMESTAMP,
					Transform: transform.FromField("Description.CompletedAt").Transform(transform.UnixToTimestamp),
					Description: "Timestamp of when the model was completed.",
				},
				{
					Name : "status",
					Type : proto.ColumnType_STRING,
					Transform: transform.FromField("Description.Status"),
					Description: "Status of the model.",
				},
				{
					Name : "last_used",
					Type : proto.ColumnType_TIMESTAMP,
					Transform: transform.FromField("Description.LastUsed").Transform(transform.UnixToTimestamp),
					Description: "Timestamp of when the model was last used.",
				},
				{
					Name : "creator_id",
					Type : proto.ColumnType_STRING,
					Transform: transform.FromField("Description.CreatorID"),
					Description: "ID of the creator of the model.",
				},
				{
					Name : "organization_id",
					Type : proto.ColumnType_STRING,
					Transform: transform.FromField("Description.OrganizationID"),
					Description: "ID of the organization of the model.",
				},
		
			
		}),
	}
}
