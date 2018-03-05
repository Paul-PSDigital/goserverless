package goserverless_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	//. "github.com/onsi/gomega/gstruct"
	. "github.com/thepauleh/goserverless"
	. "github.com/thepauleh/goserverless/serverless"
)

var _ = Describe("GoServerless", func() {

	Context("Generate a basic yaml file matching test for the service definition", func() {

		serviceConfig := NewTemplate("myService")
		template, err := serviceConfig.YAML()

		It("should successfully return the serverless template", func() {
			Expect(err).To(BeNil())
			Expect(template).ShouldNot(BeNil())
		})

		importedTemplate, err := goserverless.Open("test/yaml/basic-service.yaml")

		It("should be the same as the imported basic file", func() {
			Expect(err).To(BeNil())
			Expect(template).To(Equal(importedTemplate))
		})
	})
})
