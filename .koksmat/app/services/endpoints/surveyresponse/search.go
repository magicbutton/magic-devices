/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.search.v2
package surveyresponse

import (
    "log"

    "github.com/magicbutton/magic-devices/applogic"
    "github.com/magicbutton/magic-devices/database"
    "github.com/magicbutton/magic-devices/services/models/surveyresponsemodel"
    . "github.com/magicbutton/magic-devices/utils"
)

func SurveyResponseSearch(query string) (*Page[surveyresponsemodel.SurveyResponse], error) {
    log.Println("Calling SurveyResponsesearch")

    return applogic.Search[database.SurveyResponse, surveyresponsemodel.SurveyResponse]("name", query, applogic.MapSurveyResponseOutgoing)

}
