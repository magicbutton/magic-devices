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

func PersonWithSurveys(id int) (*personmodel.Person, error) {
	log.Println("Calling PersonRead")

	person, err := applogic.Read[database.Person, personmodel.Person](id, applogic.MapPersonOutgoing)
	if err != nil {
		return nil, err
	}
	columns := []string{"respondent_id", "survey_id"}
	survey_responses, err := applogic.SelectDistinct[database.SurveyResponse, surveyresponsemodel.SurveyResponse]("respondent_id = ?", applogic.MapSurveyResponseOutgoing, columns, person.ID)
	if err != nil {
		return nil, err
	}
	person.SurveyResponses = *survey_responses
	return person, nil

}
