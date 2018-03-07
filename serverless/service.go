package serverless

// Service is the Serverless Service definition
type Service struct {
	Name         string `json:"name,omitempty"`
	AwsKmsKeyArn string `json:"awsKmsKeyArn,omitempty"`
}
