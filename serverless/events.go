package serverless

// Events definition
type Events struct {
	HTTPEvent            *HTTPEvent            `json:"http,omitempty"`
	S3Event              *S3Event              `json:"s3,omitempty"`
	ScheduleEvent        *ScheduleEvent        `json:"schedule,omitempty"`
	SNSEvent             *SNSEvent             `json:"sns,omitempty"`
	SQSEvent             *SQSEvent             `json:"sqs,omitempty"`
	StreamEvent          *StreamEvent          `json:"stream,omitempty"`
	AlexaSkillEvent      *AlexaEvent           `json:"alexaSkill,omitempty"`
	AlexaSmartHomeEvent  *AlexaEvent           `json:"alexaSmartHome,omitempty"`
	IOTEvent             *IOTEvent             `json:"iot,omitempty"`
	CloudwatchEvent      *CloudwatchEvent      `json:"cloudwatchEvent,omitempty"`
	CloudwatchLogEvent   *CloudwatchLogEvent   `json:"cloudwatchLog,omitempty"`
	CognitoUserPoolEvent *CognitoUserPoolEvent `json:"cognitoUserPool,omitempty"`
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
	Bucket string      `json:"bucket,omitempty"`
	Event  string      `json:"event,omitempty"`
	Rules  interface{} `json:"rules,omitempty"`
}

// ScheduleEvent definition
type ScheduleEvent struct {
	Rate    string                 `json:"rate,omitempty"`
	Enabled bool                   `json:"enabled,omitempty"`
	Input   map[string]interface{} `json:"input,omitempty"`
}

// SNSEvent definition
type SNSEvent struct {
	TopicName   string `json:"topicName,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
}

// SQSEvent definition
type SQSEvent struct {
	ARN       string `json:"arn,omitempty"`
	BatchSize int    `json:"batchSize,omitempty"`
}

// StreamEvent definition
type StreamEvent struct {
	ARN              string `json:"arn,omitempty"`
	BatchSize        int    `json:"batchSize,omitempty"`
	StartingPosition string `json:"startingPosition,omitempty"`
	Enabled          bool   `json:"enabled,omitempty"`
}

// AlexaEvent defines a AlexaSkillEvent or AlexaSmartHomeEvent
type AlexaEvent struct {
	AppID   string `json:"appId,omitempty"`
	Enabled bool   `json:"enabled,omitempty"`
}

// IOTEvent definition
type IOTEvent struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Enabled     bool   `json:"enabled,omitempty"`
	SQL         string `json:"sql,omitempty"`
	SQLVersion  string `json:"sqlVersion,omitempty"`
}

// CloudwatchEvent definition
type CloudwatchEvent struct {
	Event     interface{}            `json:"event,omitempty"`
	Input     map[string]interface{} `json:"input,omitempty"`
	InputPath string                 `json:"inputPath,omitempty"`
}

// CloudwatchLogEvent definition
type CloudwatchLogEvent struct {
	LogGroup string `json:"logGroup,omitempty"`
	Filter   string `json:"filter,omitempty"`
}

// CognitoUserPoolEvent definition
type CognitoUserPoolEvent struct {
	Pool    string `json:"pool,omitempty"`
	Trigger string `json:"trigger,omitempty"`
}
