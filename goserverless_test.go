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

		It("should get the imported template", func() {
			Expect(err).To(BeNil())
		})
		// Start Testing the function events!
		usersCreateFunction := importedTemplate.Functions["usersCreate"]

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

		It("should read the provider vpc settings", func() {
			Expect(importedTemplate.Provider.VPC.SecurityGroupIds[0]).To(Equal("securityGroupId1"))
		})

		It("should read the stack tags", func() {
			Expect(importedTemplate.Provider.StackTags["key"]).To(Equal("value"))
		})

		It("should read the apiKeys", func() {
			Expect(importedTemplate.Provider.APIKeys[0]).To(Equal("myFirstKey"))
		})

		It("should read the iamRole", func() {
			Expect(importedTemplate.Provider.IAMRoleStatements[0]["Effect"]).To(Equal("Allow"))
		})

		It("should read the stack policy", func() {
			Expect(importedTemplate.Provider.StackPolicy[0]["Effect"]).To(Equal("Allow"))
		})

		It("should read the notificationArns", func() {
			Expect(importedTemplate.Provider.NotificationARNs[0]).To(Equal("arn:aws:sns:us-east-1:XXXXXX:mytopic"))
		})

		It("should read package params", func() {
			// Test the package
			Expect(importedTemplate.Package.Include[0]).To(Equal("src/**"))
			Expect(importedTemplate.Package.Include[1]).To(Equal("handler.js"))

			Expect(importedTemplate.Package.ExcludeDevDependencies).To(Equal(false))
		})

		It("should read functions params", func() {
			// Test the function
			Expect(usersCreateFunction.Handler).To(Equal("users.create"))
			Expect(usersCreateFunction.Name).To(Equal("${self:provider.stage}-lambdaName"))

			Expect(usersCreateFunction.MemorySize).To(Equal(512))
		})

		It("should read the functons vpc settings", func() {
			Expect(usersCreateFunction.VPC.SecurityGroupIds[0]).To(Equal("securityGroupId1"))
		})

		It("should read the http request", func() {
			// HTTP Request
			event := usersCreateFunction.Events[0].HTTPEvent
			Expect(event.Path).To(Equal("users/create"))
			Expect(event.Method).To(Equal("get"))

			Expect(event.Authorizer["name"]).To(Equal("authorizerFunc"))
			Expect(event.Authorizer["identityValidationExpression"]).To(Equal("someRegex"))
		})

		It("should read the s3 event", func() {
			// S3 Object
			event := usersCreateFunction.Events[1].S3Event
			Expect(event.Bucket).To(Equal("photos"))
			Expect(event.Event).To(Equal("s3:ObjectCreated:*"))

			// TODO: Rules
		})

		It("should read the schedule event", func() {
			// Schedule/Cron Event
			event := usersCreateFunction.Events[2].ScheduleEvent
			Expect(event.Rate).To(Equal("rate(10 minutes)"))
			Expect(event.Enabled).To(Equal(false))

			Expect(event.Input["key1"]).To(Equal("value1"))
			Expect(event.Input["key2"]).To(Equal("value2"))

			testStageParams := map[string]interface{}{
				"stage": "dev",
			}
			Expect(event.Input["stageParams"]).To(Equal(testStageParams))
		})

		It("should read the sns event", func() {
			event := usersCreateFunction.Events[3].SNSEvent
			Expect(event.TopicName).To(Equal("aggregate"))
			Expect(event.DisplayName).To(Equal("Data aggregation pipeline"))
		})

		It("should read the stream event", func() {
			event := usersCreateFunction.Events[4].StreamEvent
			Expect(event.ARN).To(Equal("arn:aws:kinesis:region:XXXXXX:stream/foo"))
			Expect(event.BatchSize).To(Equal(100))
			Expect(event.StartingPosition).To(Equal("LATEST"))
			Expect(event.Enabled).To(Equal(false))
		})

		// AlexaSkill
		It("should read the alex skill event", func() {
			event := usersCreateFunction.Events[5].AlexaSkillEvent
			Expect(event.AppID).To(Equal("amzn1.ask.skill.xx-xx-xx-xx"))
			Expect(event.Enabled).To(Equal(true))
		})

		// AlexaSmartHome
		It("should read the alex smart home event", func() {
			event := usersCreateFunction.Events[6].AlexaSmartHomeEvent
			Expect(event.AppID).To(Equal("amzn1.ask.skill.xx-xx-xx-xx"))
			Expect(event.Enabled).To(Equal(true))
		})

		// IOT
		It("should read the iot event", func() {
			event := usersCreateFunction.Events[7].IOTEvent
			Expect(event.Name).To(Equal("myIoTEvent"))
			Expect(event.Description).To(Equal("An IoT event"))
			Expect(event.Enabled).To(Equal(true))
			Expect(event.SQL).To(Equal("SELECT * FROM 'some_topic'"))
			Expect(event.SQLVersion).To(Equal("beta"))
		})

		// Cloudwatch Event
		It("should read the cloudwatch event", func() {
			event := usersCreateFunction.Events[8].CloudwatchEvent
			Expect(event.Input["key1"]).To(Equal("value1"))

			// TODO: Input StageParams cannot be extracted under current interface
			Expect(event.InputPath).To(Equal("$.stageVariables"))
		})

		// Cloudwatch Log
		It("should read the cloudwatch log event", func() {
			event := usersCreateFunction.Events[9].CloudwatchLogEvent
			Expect(event.LogGroup).To(Equal("/aws/lambda/hello"))
			Expect(event.Filter).To(Equal("{$.userIdentity.type = Root}"))
		})

		// Cognito User Pool
		It("should read the cognito user pool event", func() {
			event := usersCreateFunction.Events[10].CognitoUserPoolEvent
			Expect(event.Pool).To(Equal("MyUserPool"))
			Expect(event.Trigger).To(Equal("PreSignUp"))
		})

		It("should read resources params", func() {
			// Test the provider
			Expect(importedTemplate.Resources.Resources).ShouldNot(BeNil())
		})

		It("Should pull out the resource policy", func() {
			Expect(importedTemplate.Provider.ResourcePolicy).ShouldNot(BeNil())
		})
	})
})
