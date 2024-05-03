/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: true
---
*/
//generator:  noma3
package person

// noma2
import (
	"log"

	"github.com/magicbutton/magic-devices/applogic"
	"github.com/magicbutton/magic-devices/database"
	"github.com/magicbutton/magic-devices/services/models/personmodel"
	"github.com/magicbutton/magic-devices/services/models/surveyresponsemodel"
)

func PersonWithSurveyResponses(id int, surveyid int) (*personmodel.Person, error) {
	log.Println("Calling PersonRead")

	person, err := applogic.Read[database.Person, personmodel.Person](id, applogic.MapPersonOutgoing)
	if err != nil {
		return nil, err
	}
	survey_responses, err := applogic.Select[database.SurveyResponse, surveyresponsemodel.SurveyResponse]("respondent_id = ? AND survey_id = ?", applogic.MapSurveyResponseOutgoing, person.ID, surveyid)
	if err != nil {
		return nil, err
	}
	person.SurveyResponses = *survey_responses
	return person, nil

}
