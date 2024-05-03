/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/   
//version: pølsevogn1
package database

import (
	"time"
            "github.com/magicbutton/magic-devices/database/databasetypes"

	"github.com/uptrace/bun"
)

type SurveyRespondent struct {
	bun.BaseModel `bun:"table:surveyrespondent,alias:surveyrespondent"`

	ID             int     `bun:"id,pk,autoincrement"`
	CreatedAt      time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt      time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	DeletedAt      time.Time `bun:",soft_delete,nullzero"`
        Tenant string `json:"tenant"`
    Name string `json:"name"`
    Description string `json:"description"`
    Url string `json:"url"`
    Key string `json:"key"`
    Displayname string `json:"displayname"`
    Person databasetypes.Reference `json:"person"`
    Campaign databasetypes.Reference `json:"campaign"`
    Survey databasetypes.Reference `json:"survey"`

}

