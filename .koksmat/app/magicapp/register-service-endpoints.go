/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/
package magicapp

import (
	"github.com/magicbutton/magic-devices/services"
	"github.com/nats-io/nats.go/micro"
)

func RegisterServiceEndpoints(root micro.Group) {
	root.AddEndpoint("app", micro.HandlerFunc(services.HandleAppRequests))

	root.AddEndpoint("person", micro.HandlerFunc(services.HandlePersonRequests))
	root.AddEndpoint("application", micro.HandlerFunc(services.HandleApplicationRequests))
	root.AddEndpoint("campaign", micro.HandlerFunc(services.HandleCampaignRequests))
	root.AddEndpoint("survey", micro.HandlerFunc(services.HandleSurveyRequests))
	root.AddEndpoint("surveyresponse", micro.HandlerFunc(services.HandleSurveyResponseRequests))
}
