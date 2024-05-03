/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/
//GenerateMapModel v1.1
package applogic
import (
	//"encoding/json"
	//"time"
	"github.com/magicbutton/magic-devices/database"
	"github.com/magicbutton/magic-devices/services/models/surveymodel"
   
)


func MapSurveyOutgoing(db database.Survey) surveymodel.Survey {
    return surveymodel.Survey{
        ID:        db.ID,
        CreatedAt: db.CreatedAt,
        UpdatedAt: db.UpdatedAt,
                Tenant : db.Tenant,
        Name : db.Name,
        Description : db.Description,
        Url : db.Url,
        Key : db.Key,
        Displayname : db.Displayname,
                Owner_id : db.Owner_id,
                Campaign_id : db.Campaign_id,
        Question1 : db.Question1,
        Question2 : db.Question2,
        Question3 : db.Question3,
        Question4 : db.Question4,
        Question5 : db.Question5,
        Question6 : db.Question6,
        Question7 : db.Question7,
        Question8 : db.Question8,
        Question9 : db.Question9,
        Truefalse1 : db.Truefalse1,
        Truefalse2 : db.Truefalse2,
        Truefalse3 : db.Truefalse3,
        Datetime1 : db.Datetime1,
        Datetime2 : db.Datetime2,
        Datetime3 : db.Datetime3,
        Number1 : db.Number1,
        Number2 : db.Number2,
        Number3 : db.Number3,

    }
}

func MapSurveyIncoming(in surveymodel.Survey) database.Survey {
    return database.Survey{
        ID:        in.ID,
        CreatedAt: in.CreatedAt,
        UpdatedAt: in.UpdatedAt,
                Tenant : in.Tenant,
        Name : in.Name,
        Description : in.Description,
        Url : in.Url,
        Key : in.Key,
        Displayname : in.Displayname,
                Owner_id : in.Owner_id,
                Campaign_id : in.Campaign_id,
        Question1 : in.Question1,
        Question2 : in.Question2,
        Question3 : in.Question3,
        Question4 : in.Question4,
        Question5 : in.Question5,
        Question6 : in.Question6,
        Question7 : in.Question7,
        Question8 : in.Question8,
        Question9 : in.Question9,
        Truefalse1 : in.Truefalse1,
        Truefalse2 : in.Truefalse2,
        Truefalse3 : in.Truefalse3,
        Datetime1 : in.Datetime1,
        Datetime2 : in.Datetime2,
        Datetime3 : in.Datetime3,
        Number1 : in.Number1,
        Number2 : in.Number2,
        Number3 : in.Number3,

    }
}
