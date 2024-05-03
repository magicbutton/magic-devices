/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/   
//GenerateGoModel v1
package surveyresponsemodel
import (
	"encoding/json"
	"time"
    // "github.com/magicbutton/magic-devices/database/databasetypes"
)

func UnmarshalSurveyResponse(data []byte) (SurveyResponse, error) {
	var r SurveyResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SurveyResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type SurveyResponse struct {
    ID        int    `json:"id"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
        Tenant string `json:"tenant"`
    Name string `json:"name"`
    Description string `json:"description"`
    Url string `json:"url"`
    Responsedate time.Time `json:"responsedate"`
    Key string `json:"key"`
    Displayname string `json:"displayname"`
    Respondent_id int `json:"respondent_id"`
    Survey_id int `json:"survey_id"`
    Application_id int `json:"application_id"`
    Questions interface{} `json:"questions"`
    Answers interface{} `json:"answers"`
    Answer1 string `json:"answer1"`
    Answer2 string `json:"answer2"`
    Answer3 string `json:"answer3"`
    Answer4 string `json:"answer4"`
    Answer5 string `json:"answer5"`
    Answer6 string `json:"answer6"`
    Answer7 string `json:"answer7"`
    Answer8 string `json:"answer8"`
    Answer9 string `json:"answer9"`
    Truefalse1 bool `json:"truefalse1"`
    Truefalse2 bool `json:"truefalse2"`
    Truefalse3 bool `json:"truefalse3"`
    Datetime1 time.Time `json:"datetime1"`
    Datetime2 time.Time `json:"datetime2"`
    Datetime3 time.Time `json:"datetime3"`
    Number1 int `json:"number1"`
    Number2 int `json:"number2"`
    Number3 int `json:"number3"`

}

