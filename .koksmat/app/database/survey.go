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

type Survey struct {
	bun.BaseModel `bun:"table:survey,alias:survey"`

	ID             int     `bun:"id,pk,autoincrement"`
	CreatedAt      time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt      time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	DeletedAt      time.Time `bun:",soft_delete,nullzero"`
        Tenant string `bun:"tenant"`
    Name string `bun:"name"`
    Description string `bun:"description"`
    Url string `bun:"url"`
    Key string `bun:"key"`
    Displayname string `bun:"displayname"`
    Owner_id int `bun:"owner_id"`
    Campaign_id int `bun:"campaign_id"`
    Question1 string `bun:"question1"`
    Question2 string `bun:"question2"`
    Question3 string `bun:"question3"`
    Question4 string `bun:"question4"`
    Question5 string `bun:"question5"`
    Question6 string `bun:"question6"`
    Question7 string `bun:"question7"`
    Question8 string `bun:"question8"`
    Question9 string `bun:"question9"`
    Truefalse1 string `bun:"truefalse1"`
    Truefalse2 string `bun:"truefalse2"`
    Truefalse3 string `bun:"truefalse3"`
    Datetime1 string `bun:"datetime1"`
    Datetime2 string `bun:"datetime2"`
    Datetime3 string `bun:"datetime3"`
    Number1 string `bun:"number1"`
    Number2 string `bun:"number2"`
    Number3 string `bun:"number3"`
    Questions interface{} `bun:"questions"`

}

