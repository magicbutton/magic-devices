/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/   
//version: pølsevogn2
package database

import (
	"time"
    
	"github.com/uptrace/bun"
)

type SurveyResponse struct {
	bun.BaseModel `bun:"table:surveyresponse,alias:surveyresponse"`

	ID             int     `bun:"id,pk,autoincrement"`
	CreatedAt      time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt      time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	DeletedAt      time.Time `bun:",soft_delete,nullzero"`
        Tenant string `bun:"tenant"`
    Name string `bun:"name"`
    Description string `bun:"description"`
    Url string `bun:"url"`
    Responsedate time.Time `bun:"responsedate"`
    Key string `bun:"key"`
    Displayname string `bun:"displayname"`
    Respondent_id int `bun:"respondent_id"`
    Survey_id int `bun:"survey_id"`
    Application_id int `bun:"application_id"`
    Questions interface{} `bun:"questions"`
    Answers interface{} `bun:"answers"`
    Answer1 string `bun:"answer1"`
    Answer2 string `bun:"answer2"`
    Answer3 string `bun:"answer3"`
    Answer4 string `bun:"answer4"`
    Answer5 string `bun:"answer5"`
    Answer6 string `bun:"answer6"`
    Answer7 string `bun:"answer7"`
    Answer8 string `bun:"answer8"`
    Answer9 string `bun:"answer9"`
    Truefalse1 bool `bun:"truefalse1"`
    Truefalse2 bool `bun:"truefalse2"`
    Truefalse3 bool `bun:"truefalse3"`
    Datetime1 time.Time `bun:"datetime1"`
    Datetime2 time.Time `bun:"datetime2"`
    Datetime3 time.Time `bun:"datetime3"`
    Number1 int `bun:"number1"`
    Number2 int `bun:"number2"`
    Number3 int `bun:"number3"`

}

