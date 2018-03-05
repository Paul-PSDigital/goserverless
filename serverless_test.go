package serverless_test

import (
	. "github.com/thepauleh/goserverless/serverless"
)

var _ = Describe("GoServerless", func() {

	Context("Generate a basic yaml file matching test for the service definition", func() {

		serviceConfig := NewTemplate("myService")
		data, err := serviceConfig.YAML()
	}
}
