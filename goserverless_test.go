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

		// Start Testing the function events!
		usersCreateFunction := importedTemplate.Functions["usersCreate"]

		It("should read functions params", func() {
			// Test the function
			Expect(usersCreateFunction.Handler).To(Equal("users.create"))
			Expect(usersCreateFunction.Name).To(Equal("${self:provider.stage}-lambdaName"))

			Expect(usersCreateFunction.MemorySize).To(Equal(512))
		})

		It("should read the http request", func() {
			// HTTP Request
			httpEvent := usersCreateFunction.Events[0].HTTPEvent
			Expect(httpEvent.Path).To(Equal("users/create"))
			Expect(httpEvent.Method).To(Equal("get"))

			Expect(httpEvent.Authorizer["name"]).To(Equal("authorizerFunc"))
			Expect(httpEvent.Authorizer["identityValidationExpression"]).To(Equal("someRegex"))
		})

		It("should read the s3 event", func() {
			// S3 Object
			s3Event := usersCreateFunction.Events[1].S3Event
			Expect(s3Event.Bucket).To(Equal("photos"))
			Expect(s3Event.Event).To(Equal("s3:ObjectCreated:*"))

			// TODO: Rules
		})

		It("should read the schedule event", func() {
			// Schedule/Cron Event
			httpEvent := usersCreateFunction.Events[2].ScheduleEvent
			Expect(httpEvent.Rate).To(Equal("rate(10 minutes)"))
			Expect(httpEvent.Enabled).To(Equal(false))

			Expect(httpEvent.Input["key1"]).To(Equal("value1"))
			Expect(httpEvent.Input["key2"]).To(Equal("value2"))

			testStageParams := map[string]interface{}{
				"stage": "dev",
			}
			Expect(httpEvent.Input["stageParams"]).To(Equal(testStageParams))
		})

		// TODO:
		// SNS
		// Stream
		// AlexaSkill
		// AlexaSmartHome
		// IOT
		// Cloudwatch Event
		// Cloudwatch Log
		// Cognito User Pool

		It("should read resources params", func() {
			// Test the provider
			Expect(importedTemplate.Resources.Resources).ShouldNot(BeNil())
		})
	})
})
