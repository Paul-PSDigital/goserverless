package serverless

import (
	"github.com/awslabs/goformation/cloudformation"
	"github.com/sanathkr/yaml"
)

// Template represents an AWS CloudFormation template
// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/template-anatomy.html
type Template struct {
	Service          *Service                `json:"service,omitempty"` // SHOULD BE TYPE SERVICE
	FrameworkVersion string                  `json:"frameworkVersion,omitempty"`
	Functions        map[string]interface{}  `json:"functions,omitempty"`
	Resources        cloudformation.Template `json:"resources,omitempty"`
	Provider         string                  `json:"provider,omitempty"` // SHOULD BE TYPE PROVIDER
	Package          string                  `json:"pacakge,omitempty"`  // SHOULD BE THE PACKAGE TYPE
}

// NewTemplate creates a new AWS CloudFormation template struct
func NewTemplate(serviceName string) *Template {
	return &Template{
		Service:          &Service{Name: serviceName},
		FrameworkVersion: ">=1.0.0 <2.0.0",
	}
}

// YAML converts an AWS CloudFormation template object to YAML
func (t *Template) YAML() ([]byte, error) {
	return yaml.Marshal(t)
}
