package goserverless

import (
	"io/ioutil"

	"github.com/thepauleh/goserverless/serverless"
	yaml "gopkg.in/yaml.v2"
)

// Open and parse a Serverless template from file.
// Works with YAML formatted templates.
func Open(filename string) (*serverless.Template, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	template := &serverless.Template{}
	if err := yaml.Unmarshal(data, template); err != nil {
		return nil, err
	}

	return template, nil
}
