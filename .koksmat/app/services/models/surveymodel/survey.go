/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/   
//GenerateGoModel v1
package surveymodel
import (
	"encoding/json"
	"time"
    // "github.com/magicbutton/magic-devices/database/databasetypes"
)

func UnmarshalSurvey(data []byte) (Survey, error) {
	var r Survey
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Survey) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Survey struct {
    ID        int    `json:"id"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
        Tenant string `json:"tenant"`
    Name string `json:"name"`
    Description string `json:"description"`
    Url string `json:"url"`
    Key string `json:"key"`
    Displayname string `json:"displayname"`
    Owner_id int `json:"owner_id"`
    Campaign_id int `json:"campaign_id"`
    Question1 string `json:"question1"`
    Question2 string `json:"question2"`
    Question3 string `json:"question3"`
    Question4 string `json:"question4"`
    Question5 string `json:"question5"`
    Question6 string `json:"question6"`
    Question7 string `json:"question7"`
    Question8 string `json:"question8"`
    Question9 string `json:"question9"`
    Truefalse1 string `json:"truefalse1"`
    Truefalse2 string `json:"truefalse2"`
    Truefalse3 string `json:"truefalse3"`
    Datetime1 string `json:"datetime1"`
    Datetime2 string `json:"datetime2"`
    Datetime3 string `json:"datetime3"`
    Number1 string `json:"number1"`
    Number2 string `json:"number2"`
    Number3 string `json:"number3"`
    Questions interface{} `json:"questions"`

}

