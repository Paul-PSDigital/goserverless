package goserverless_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/thepauleh/goserverless"
)

var _ = Describe("GoServerless", func() {

	Context("Transform and hydrate a basic serverless config back to our structs", func() {
		importedTemplate, err := goserverless.Open("test/yaml/basic-service.yaml")

		It("should read basic params", func() {
			Expect(err).To(BeNil())

			Expect(importedTemplate.Service.Name).To(Equal("basicService"))
			Expect(importedTemplate.FrameworkVersion).To(Equal(">=1.0.0 <2.1.9"))
		})
	})

	Context("Transform and hydrate a full serverless config back to our structs", func() {
		importedTemplate, err := goserverless.Open("test/yaml/simple-serverless.yaml")

		It("should read basic params", func() {
			Expect(err).To(BeNil())

			Expect(importedTemplate.Service.Name).To(Equal("myService"))
			Expect(importedTemplate.FrameworkVersion).To(Equal(">=1.0.0 <2.0.0"))
		})

		It("should read provider params", func() {
			// Test the provider
			Expect(importedTemplate.Provider.Name).To(Equal("aws"))
			Expect(importedTemplate.Provider.Runtime).To(Equal("nodejs6.10"))
			Expect(importedTemplate.Provider.MemorySize).To(Equal(512))

			const serviceEnvVar float64 = 123456789
			Expect(importedTemplate.Provider.Environment["serviceEnvVar"]).To(Equal(serviceEnvVar))
		})

		It("should read package params", func() {
			// Test the package
			Expect(importedTemplate.Package.Include[0]).To(Equal("src/**"))
			Expect(importedTemplate.Package.Include[1]).To(Equal("handler.js"))

			Expect(importedTemplate.Package.ExcludeDevDependencies).To(Equal(false))
		})

		It("should read functions params", func() {
			usersCreateFunction := importedTemplate.Functions["usersCreate"]
			// Test the function
			Expect(usersCreateFunction.Handler).To(Equal("users.create"))
			Expect(usersCreateFunction.Name).To(Equal("${self:provider.stage}-lambdaName"))

			Expect(usersCreateFunction.MemorySize).To(Equal(512))
			// TODO: Events
		})

		It("should read resources params", func() {
			// Test the provider
			Expect(importedTemplate.Resources.Resources).ShouldNot(BeNil())
		})
	})
})
