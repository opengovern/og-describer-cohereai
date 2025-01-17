//go:generate go run ../../pkg/sdk/runable/steampipe_es_client_generator/main.go -pluginPath ../../steampipe-plugin-REPLACEME/REPLACEME -file $GOFILE -output ../../pkg/sdk/es/resources_clients.go -resourceTypesFile ../resource_types/resource-types.json

// Implement types for each resource
package provider
import "time"

type Metadata struct{}

type ConnectorDescription struct {
	Connectors []Connector `json:"connectors"`
	TotalCount float64     `json:"total_count"`
}
type ConnectorDetailResponse struct {
	Connector Connector `json:"connector"`
}


type Connector struct {
	ID                string    `json:"id"`
	Name              string    `json:"name"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	OrganizationID    string    `json:"organization_id"`
	Description       string    `json:"description"`
	URL               string    `json:"url"`
	Excludes          []string  `json:"excludes"`
	AuthType          string    `json:"auth_type"`
	Oauth             Oauth     `json:"oauth"`
	AuthStatus        string    `json:"auth_status"`
	Active            bool      `json:"active"`
	ContinueOnFailure bool      `json:"continue_on_failure"`
}

type Oauth struct {
	AuthorizeURL string `json:"authorize_url"`
	TokenURL     string `json:"token_url"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Scope        string `json:"scope"`
}

type ModelListResponse struct {
	Models        []ModelDescription `json:"models"`
	NextPageToken string  `json:"next_page_token"`
}

type ModelDescription struct {
	Name             string   `json:"name"`
	Endpoints        []string `json:"endpoints"`
	Finetuned        bool     `json:"finetuned"`
	ContextLength    float64  `json:"context_length"`
	TokenizerURL     string   `json:"tokenizer_url"`
	DefaultEndpoints []string `json:"default_endpoints"`
}

type DatasetListResponse struct {
	Datasets []DatasetDescription `json:"datasets"`
}


type DatasetDescription struct {
	ID                 string        `json:"id"`
	Name               string        `json:"name"`
	CreatedAt          time.Time     `json:"created_at"`
	UpdatedAt          time.Time     `json:"updated_at"`
	DatasetType        string        `json:"dataset_type"`
	ValidationStatus   string        `json:"validation_status"`
	ValidationError    string        `json:"validation_error"`
	Schema             string        `json:"schema"`
	RequiredFields     []string      `json:"required_fields"`
	PreserveFields     []string      `json:"preserve_fields"`
	DatasetParts       []DatasetPart `json:"dataset_parts"`
	ValidationWarnings []string      `json:"validation_warnings"`
	TotalUsage         float64       `json:"total_usage"`
}

type DatasetPart struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type OrganizationUsage struct {
	OrganizationUsage int64 `json:"organization_usage"`
}


type ListFineTunedModelsResponse struct {
	FinetunedModels []FineTunedModelDescription `json:"finetuned_models"`
	NextPageToken   string           `json:"next_page_token"`
	TotalSize       int64            `json:"total_size"`
}



type FineTunedModelDescription struct {
	Name           string    `json:"name"`
	Settings       Settings  `json:"settings"`
	ID             string    `json:"id"`
	CreatorID      string    `json:"creator_id"`
	OrganizationID string    `json:"organization_id"`
	Status         string    `json:"status"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	CompletedAt    time.Time `json:"completed_at"`
	LastUsed       time.Time `json:"last_used"`
}

type Settings struct {
	BaseModel       BaseModel        `json:"base_model"`
	DatasetID       string           `json:"dataset_id"`
	Hyperparameters *Hyperparameters `json:"hyperparameters,omitempty"`
}

type BaseModel struct {
	BaseType string `json:"base_type"`
	Name     string `json:"name"`
	Version  string `json:"version"`
	Strategy string `json:"strategy"`
}

type Hyperparameters struct {
	EarlyStoppingPatience  int64   `json:"early_stopping_patience"`
	EarlyStoppingThreshold float64 `json:"early_stopping_threshold"`
	TrainBatchSize         int64   `json:"train_batch_size"`
	TrainEpochs            int64   `json:"train_epochs"`
	LearningRate           float64 `json:"learning_rate"`
}


type ListEmbedJobsResponse struct {
	EmbedJobs []EmbedJobDescription `json:"embed_jobs"`
}


type EmbedJobDescription struct {
	JobID           string    `json:"job_id"`
	Status          string    `json:"status"`
	CreatedAt       time.Time `json:"created_at"`
	InputDatasetID  string    `json:"input_dataset_id"`
	Model           string    `json:"model"`
	Truncate        string    `json:"truncate"`
	Name            string    `json:"name"`
	OutputDatasetID string    `json:"output_dataset_id"`
}
