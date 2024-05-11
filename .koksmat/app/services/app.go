/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: true
---
*/
// macd.1
package services

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/nats-io/nats.go/micro"

	"github.com/magicbutton/magic-devices/services/endpoints/app"
	. "github.com/magicbutton/magic-devices/utils"
)

func HandleAppRequests(req micro.Request) {

	rawRequest := string(req.Data())
	if rawRequest == "ping" {
		req.Respond([]byte("pong"))
		return

	}

	var payload ServiceRequest
	_ = json.Unmarshal([]byte(req.Data()), &payload)
	if len(payload.Args) < 1 {
		ServiceResponseError(req, "missing command")
		return

	}
	switch payload.Args[0] {

	// macd.2
	case "dashboard":
		if len(payload.Args) < 2 {
			log.Println("Expected 2 arguments, got %d", len(payload.Args))
			ServiceResponseError(req, "Expected 1 arguments")
			return
		}

		result, err := app.GlobalDashboard(payload.Args[1])
		if err != nil {
			log.Println("Error", err)
			ServiceResponseError(req, fmt.Sprintf("Error calling GlobalDashboard: %s", err))

			return
		}

		ServiceResponse(req, result)

	// macd.2

	default:
		ServiceResponseError(req, "Unknown command")
	}
}
