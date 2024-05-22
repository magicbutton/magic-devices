/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/
// macd.1
package services

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/nats-io/nats.go/micro"

	"github.com/magicbutton/magic-devices/services/endpoints/exceptiontype"
	"github.com/magicbutton/magic-devices/services/models/exceptiontypemodel"
	. "github.com/magicbutton/magic-devices/utils"
)

func HandleExceptiontypeRequests(req micro.Request) {

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
	case "read":
		if len(payload.Args) < 2 {
			log.Println("Expected 2 arguments, got %d", len(payload.Args))
			ServiceResponseError(req, "Expected 1 arguments")
			return
		}

		result, err := exceptiontype.ExceptiontypeRead(StrToInt(payload.Args[1]))
		if err != nil {
			log.Println("Error", err)
			ServiceResponseError(req, fmt.Sprintf("Error calling ExceptiontypeRead: %s", err))

			return
		}

		ServiceResponse(req, result)

	// macd.2
	case "create":
		if len(payload.Args) < 2 {
			log.Println("Expected 2 arguments, got %d", len(payload.Args))
			ServiceResponseError(req, "Expected 1 arguments")
			return
		}

		// transformer v1
		object := exceptiontypemodel.Exceptiontype{}
		body := ""

		json.Unmarshal([]byte(payload.Args[1]), &body)
		err := json.Unmarshal([]byte(body), &object)

		if err != nil {
			log.Println("Error", err)
			ServiceResponseError(req, "Error unmarshalling exceptiontype")
			return
		}

		result, err := exceptiontype.ExceptiontypeCreate(object)
		if err != nil {
			log.Println("Error", err)
			ServiceResponseError(req, fmt.Sprintf("Error calling ExceptiontypeCreate: %s", err))

			return
		}

		ServiceResponse(req, result)

	// macd.2
	case "update":
		if len(payload.Args) < 2 {
			log.Println("Expected 2 arguments, got %d", len(payload.Args))
			ServiceResponseError(req, "Expected 1 arguments")
			return
		}

		// transformer v1
		object := exceptiontypemodel.Exceptiontype{}
		body := ""

		json.Unmarshal([]byte(payload.Args[1]), &body)
		err := json.Unmarshal([]byte(body), &object)

		if err != nil {
			log.Println("Error", err)
			ServiceResponseError(req, "Error unmarshalling exceptiontype")
			return
		}

		result, err := exceptiontype.ExceptiontypeUpdate(object)
		if err != nil {
			log.Println("Error", err)
			ServiceResponseError(req, fmt.Sprintf("Error calling ExceptiontypeUpdate: %s", err))

			return
		}

		ServiceResponse(req, result)

	// macd.2
	case "delete":
		if len(payload.Args) < 2 {
			log.Println("Expected 2 arguments, got %d", len(payload.Args))
			ServiceResponseError(req, "Expected 1 arguments")
			return
		}

		err := exceptiontype.ExceptiontypeDelete(StrToInt(payload.Args[1]))
		if err != nil {
			log.Println("Error", err)
			ServiceResponseError(req, fmt.Sprintf("Error calling ExceptiontypeDelete: %s", err))

			return
		}
		ServiceResponse(req, "")

	// macd.2
	case "search":
		if len(payload.Args) < 2 {
			log.Println("Expected 2 arguments, got %d", len(payload.Args))
			ServiceResponseError(req, "Expected 1 arguments")
			return
		}

		result, err := exceptiontype.ExceptiontypeSearch(payload.Args[1])
		if err != nil {
			log.Println("Error", err)
			ServiceResponseError(req, fmt.Sprintf("Error calling ExceptiontypeSearch: %s", err))

			return
		}

		ServiceResponse(req, result)

	default:
		ServiceResponseError(req, "Unknown command")
	}
}
