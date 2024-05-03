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
	"time"

	"github.com/magicbutton/magic-devices/services/endpoints/surveyresponse"
	"github.com/magicbutton/magic-devices/services/models/surveyresponsemodel"

	. "github.com/magicbutton/magic-devices/utils"
	"github.com/nats-io/nats.go/micro"
)

func HandleSurveyResponseRequests(req micro.Request) {

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

		result, err := surveyresponse.SurveyResponseRead(StrToInt(payload.Args[1]))
		if err != nil {
			log.Println("Error", err)
			ServiceResponseError(req, fmt.Sprintf("Error calling SurveyResponseRead: %s", err))

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
		object := surveyresponsemodel.SurveyResponse{}
		body := ""

		json.Unmarshal([]byte(payload.Args[1]), &body)
		err := json.Unmarshal([]byte(body), &object)

		if err != nil {
			log.Println("Error", err)
			ServiceResponseError(req, "Error unmarshalling surveyresponse")
			return
		}

		result, err := surveyresponse.SurveyResponseCreate(object)
		if err != nil {
			log.Println("Error", err)
			ServiceResponseError(req, fmt.Sprintf("Error calling SurveyResponseCreate: %s", err))

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
		object := surveyresponsemodel.SurveyResponse{}
		body := ""

		json.Unmarshal([]byte(payload.Args[1]), &body)
		err := json.Unmarshal([]byte(body), &object)

		if err != nil {
			log.Println("Error", err)
			ServiceResponseError(req, "Error unmarshalling surveyresponse")
			return
		}

		result, err := surveyresponse.SurveyResponseUpdate(object)
		if err != nil {
			log.Println("Error", err)
			ServiceResponseError(req, fmt.Sprintf("Error calling SurveyResponseUpdate: %s", err))

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

		err := surveyresponse.SurveyResponseDelete(StrToInt(payload.Args[1]))
		if err != nil {
			log.Println("Error", err)
			ServiceResponseError(req, fmt.Sprintf("Error calling SurveyResponseDelete: %s", err))

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

		result, err := surveyresponse.SurveyResponseSearch(payload.Args[1])
		if err != nil {
			log.Println("Error", err)
			ServiceResponseError(req, fmt.Sprintf("Error calling SurveyResponseSearch: %s", err))

			return
		}

		ServiceResponse(req, result)
	case "submit":
		if len(payload.Args) < 2 {
			log.Println("Expected 2 arguments, got %d", len(payload.Args))
			ServiceResponseError(req, "Expected 1 arguments")
			return
		}

		// transformer v1
		object := []surveyresponsemodel.SurveyResponse{}
		//		body := ""

		err := json.Unmarshal([]byte(payload.Args[1]), &object)
		//err := json.Unmarshal([]byte(body), &object)

		if err != nil {
			log.Println("Error", err)
			ServiceResponseError(req, "Error unmarshalling surveyresponse")
			return
		}

		for _, item := range object {
			item.Responsedate = time.Now()
			item.UpdatedAt = time.Now()

			_, err := surveyresponse.SurveyResponseUpdate(item)
			if err != nil {
				log.Println("Error", err)
				ServiceResponseError(req, fmt.Sprintf("Error calling SurveyResponseUpdate: %s", err))
				return
			}
		}

		if err != nil {
			log.Println("Error", err)
			ServiceResponseError(req, fmt.Sprintf("Error calling SurveyResponseUpdate: %s", err))

			return
		}

		ServiceResponse(req, object)

	default:
		ServiceResponseError(req, "Unknown command")
	}
}
