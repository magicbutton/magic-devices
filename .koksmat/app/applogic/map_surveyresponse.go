/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: true
---
*/
//GenerateMapModel v1.1
package applogic

import (
	//"encoding/json"
	//"time"
	"github.com/magicbutton/magic-devices/database"
	"github.com/magicbutton/magic-devices/services/models/surveyresponsemodel"
)

func MapSurveyResponseOutgoing(db database.SurveyResponse) surveyresponsemodel.SurveyResponse {
	return surveyresponsemodel.SurveyResponse{
		ID:             db.ID,
		CreatedAt:      db.CreatedAt,
		UpdatedAt:      db.UpdatedAt,
		Tenant:         db.Tenant,
		Responsedate:   db.Responsedate,
		Name:           db.Name,
		Description:    db.Description,
		Url:            db.Url,
		Key:            db.Key,
		Displayname:    db.Displayname,
		Respondent_id:  db.Respondent_id,
		Survey_id:      db.Survey_id,
		Application_id: db.Application_id,
		Answer1:        db.Answer1,
		Answer2:        db.Answer2,
		Answer3:        db.Answer3,
		Answer4:        db.Answer4,
		Answer5:        db.Answer5,
		Answer6:        db.Answer6,
		Answer7:        db.Answer7,
		Answer8:        db.Answer8,
		Answer9:        db.Answer9,
	}
}

func MapSurveyResponseIncoming(in surveyresponsemodel.SurveyResponse) database.SurveyResponse {
	return database.SurveyResponse{
		ID:             in.ID,
		CreatedAt:      in.CreatedAt,
		UpdatedAt:      in.UpdatedAt,
		Tenant:         in.Tenant,
		Responsedate:   in.Responsedate,
		Name:           in.Name,
		Description:    in.Description,
		Url:            in.Url,
		Key:            in.Key,
		Displayname:    in.Displayname,
		Respondent_id:  in.Respondent_id,
		Survey_id:      in.Survey_id,
		Application_id: in.Application_id,
		Answer1:        in.Answer1,
		Answer2:        in.Answer2,
		Answer3:        in.Answer3,
		Answer4:        in.Answer4,
		Answer5:        in.Answer5,
		Answer6:        in.Answer6,
		Answer7:        in.Answer7,
		Answer8:        in.Answer8,
		Answer9:        in.Answer9,
	}
}
