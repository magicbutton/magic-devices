/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3
package app

// noma2
import (
	"context"
	"fmt"
	"log"

	"github.com/magicbutton/magic-devices/services/endpoints/person"
	"github.com/magicbutton/magic-devices/services/models"
	"github.com/magicbutton/magic-devices/utils"
	"github.com/uptrace/bun"
)

type Count struct {
	bun.BaseModel `bun:"table:jobs,alias:j"`

	Count int `json:"count"`
}

func GetCount(sql string) int {
	ctx := context.Background()
	count := []Count{}
	rows, err := utils.Db.QueryContext(ctx, sql)
	if err != nil {
		return -1
	}

	err = utils.Db.ScanRows(ctx, rows, &count)

	return count[0].Count
}

func GlobalDashboard(email string) (*models.Dashboard, error) {
	var person_id int = -1
	var numberOfAppsOwned int = 0
	var numberOfSurveyResponses int = 0
	personSearchResult, err := person.PersonSearch(("%email:" + email + "%"))
	if err == nil {
		if personSearchResult.TotalItems > 0 {
			person_id = personSearchResult.Items[0].ID
			numberOfAppsOwned = GetCount(fmt.Sprintf("SELECT COUNT(*) as count FROM public.application  where owner_id = %d ", person_id))
			numberOfSurveyResponses = GetCount(fmt.Sprintf("SELECT COUNT(*) as count FROM public.surveyresponse  where respondent_id = %d AND responsedate is null", person_id))
		}
	}

	log.Println("Calling Dashboard with person_id: ", person_id, " email: ", email)

	dashboard := models.Dashboard{
		PersonId:                person_id,
		NumberOfApplications:    GetCount("SELECT count(*) as count FROM public.application"),
		NumberOfSurveys:         GetCount("SELECT count(*) as count FROM public.survey"),
		NumberOfOwners:          GetCount("SELECT count(*) as count FROM public.person"),
		NumberOfAppsYouOwn:      numberOfAppsOwned,
		NumberOfSurveyResponses: numberOfSurveyResponses,
	}

	return &dashboard, nil

}
