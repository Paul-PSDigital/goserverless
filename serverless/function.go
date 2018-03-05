package serverless

// Function definition
type Function struct {
	Handler      string                 `json:"handler,omitempty"`
	Name         string                 `json:"name,omitempty"`
	Description  string                 `json:"description,omitempty"`
	MemorySize   int                    `json:"memorySize,omitempty"`
	Runtime      string                 `json:"runtime,omitempty"`
	Timeout      int                    `json:"timeout,omitempty"`
	Role         string                 `json:"role,omitempty"`
	OnError      string                 `json:"onError,omitempty"`
	AwsKmsKeyArn string                 `json:"awsKmsKeyArn,omitempty"`
	Environment  map[string]interface{} `json:"environment,omitempty"`
	Tags         map[string]interface{} `json:"tags,omitempty"`
	Events       []Events               `json:"events,omitempty"`
}

// HTTPEvent definition
type HTTPEvent struct {
	Path       string                 `json:"path,omitempty"`
	Method     string                 `json:"method,omitempty"`
	Cors       bool                   `json:"cors,omitempty"`
	Private    bool                   `json:"private,omitempty"`
	Authorizer map[string]interface{} `json:"authorizer,omitempty"`
}

// S3Event definition
type S3Event struct {
	Bucket string `json:"bucket,omitempty"`
	Event  string `json:"event,omitempty"`
	// TODO: Rules
}

// ScheduleEvent definition
type ScheduleEvent struct {
	Rate    string                 `json:"rate,omitempty"`
	Enabled bool                   `json:"enabled,omitempty"`
	Input   map[string]interface{} `json:"input,omitempty"`
}

// VPC definition
type VPC struct {
	SecurityGroupIds []string `json:"securityGroupIds,omitempty"`
	SubnetIds        []string `json:"subnetIds,omitempty"`
}

// Events definition
type Events struct {
	HTTPEvent     *HTTPEvent     `json:"http,omitempty"`
	S3Event       *S3Event       `json:"s3,omitempty"`
	ScheduleEvent *ScheduleEvent `json:"schedule,omitempty"`
}
