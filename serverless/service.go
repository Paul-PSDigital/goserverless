package serverless

// Service is the Serverless Service definition
type Service struct {
	Name         string `json:"name,omitempty"`         // SHOULD BE TYPE PROVIDER
	AwsKmsKeyArn string `json:"awsKmsKeyArn,omitempty"` // SHOULD BE THE PACKAGE TYPE
}
