/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/
package magicapp

import (
	"github.com/nats-io/nats.go/micro"

	"github.com/magicbutton/magic-devices/services"
)

func RegisterServiceEndpoints(root micro.Group) {
	root.AddEndpoint("person", micro.HandlerFunc(services.HandlePersonRequests))
	root.AddEndpoint("devicetype", micro.HandlerFunc(services.HandleDevicetypeRequests))
	root.AddEndpoint("device", micro.HandlerFunc(services.HandleDeviceRequests))
	root.AddEndpoint("exceptiontype", micro.HandlerFunc(services.HandleExceptiontypeRequests))
	root.AddEndpoint("grantedexception", micro.HandlerFunc(services.HandleGrantedexceptionRequests))
	root.AddEndpoint("importdata", micro.HandlerFunc(services.HandleImportdataRequests))
}
