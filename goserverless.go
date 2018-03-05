package goserverless

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	yamlwrapper "github.com/sanathkr/yaml"

	"github.com/thepauleh/goserverless/serverless"
)

// Open and parse a Serverless template from file.
// Works with YAML formatted templates.
func Open(filename string) (*serverless.Template, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return openYaml(data)
}

func openYaml(input []byte) (*serverless.Template, error) {
	data, err := yamlwrapper.YAMLToJSON(input)
	if err != nil {
		return nil, fmt.Errorf("invalid YAML template: %s", err)
	}

	template := &serverless.Template{}
	if err := json.Unmarshal(data, template); err != nil {
		return nil, err
	}

	return template, nil
}
