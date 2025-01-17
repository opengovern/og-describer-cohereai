package cohere

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
	opengovernance "github.com/opengovern/og-describer-cohereai/discovery/pkg/es"

)

func tableCohereConnectors(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "cohereai_connectors",
		Description: "Cohere Ai list of connectors.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListConnector,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    opengovernance.GetConnector,
		},
		Columns: commonColumns([]*plugin.Column{
			// Top columns
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ID"),
				Description: "ID of Connector."},
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Name"),
				Description: "Name of the connector."},
			{
				Name:        "created_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.CreatedAt").Transform(transform.UnixToTimestamp),
				Description: "Timestamp of when the connector was created."},
			{
				Name:        "updated_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.UpdatedAt").Transform(transform.UnixToTimestamp),
				Description: "Timestamp of when the connector was updated"},
			// Other columns
			{
				Name:        "organization_id",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.OrganizationId"),
				Description: "ID of the organization that owns the connector."},
			//{
			//	Name: "owner",
			//	Type: proto.ColumnType_STRING,
			//	Description: "Organization that owns the file, e.g. openai."},
			{
				Name:        "url",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Url"),
				Description: "URL of the connector."},
			
			{
				Name:        "excludes",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Excludes"),
				Description: "List of excluded fields."},
			{
				Name:        "description",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Description"),
				Description: "Description of the connector."},
			
			
		}),
	}
}
