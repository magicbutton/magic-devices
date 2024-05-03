/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3
package survey

// noma2
import (
	"log"

	"github.com/magicbutton/magic-devices/applogic"
	"github.com/magicbutton/magic-devices/database"
	"github.com/magicbutton/magic-devices/services/models/surveymodel"
)

func SurveyRead(id int) (*surveymodel.Survey, error) {
	log.Println("Calling SurveyRead")

	return applogic.Read[database.Survey, surveymodel.Survey](id, applogic.MapSurveyOutgoing)

}
