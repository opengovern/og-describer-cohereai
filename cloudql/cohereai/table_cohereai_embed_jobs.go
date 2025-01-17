package cohere

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
	opengovernance "github.com/opengovern/og-describer-cohereai/discovery/pkg/es"


)

func tableCohereEmbedJobs(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "cohereai_embed_jobs",
		Description: "Cohere Ai list of EmbedJobs.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListEmbedJob,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("job_id"),
			Hydrate:    opengovernance.GetEmbedJob,
		},
		Columns: commonColumns([]*plugin.Column{
			// Top columns
			{
				Name:        "job_id",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.job_id"),
				Description: "job_id of Embed job."},
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.name"),
				Description: "Name of the Embed job."},
			{
				Name:        "status",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.status"),
				Description: "Status of the Embed job."},
			{
				Name:        "created_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.created_at").Transform(transform.UnixToTimestamp),
				Description: "Timestamp of when the Embed job was created."},

			{
				Name:        "model",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.model"),
				Description: "Model of the Embed job."},
			{
				Name:        "input_dataset_id",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.input_dataset_id"),
				Description: "Input dataset id of the Embed job."},
			{
				Name:        "output_dataset_id",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.output_dataset_id"),
				Description: "Output dataset id of the Embed job."},
			{
				Name:        "truncate",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.truncate"),
				Description: "Truncate of the Embed job."},
			
			
		}),
	}
}
