/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: true
---
*/
//GenerateGoModel v1
package personmodel

import (
	"encoding/json"
	"time"

	"github.com/magicbutton/magic-devices/services/models/applicationmodel"
	"github.com/magicbutton/magic-devices/services/models/surveyresponsemodel"
)

func UnmarshalPerson(data []byte) (Person, error) {
	var r Person
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Person) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Person struct {
	ID               int                                  `json:"id"`
	CreatedAt        time.Time                            `json:"created_at"`
	UpdatedAt        time.Time                            `json:"updated_at"`
	Tenant           string                               `json:"tenant"`
	Name             string                               `json:"name"`
	Description      string                               `json:"description"`
	Unique_Person_Id string                               `json:"unique_person_id"`
	Displayname      string                               `json:"displayname"`
	Email            string                               `json:"email"`
	Applications     []applicationmodel.Application       `json:"applications"`
	SurveyResponses  []surveyresponsemodel.SurveyResponse `json:"survey_responses"`
}
