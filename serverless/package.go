package serverless

// Package definition
type Package struct {
	ExcludeDevDependencies bool     `json:"excludeDevDependencies,omitempty"`
	Include                []string `json:"include,omitempty"`
	Exclude                []string `json:"exclude,omitempty"`
}
