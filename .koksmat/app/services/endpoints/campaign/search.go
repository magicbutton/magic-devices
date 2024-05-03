/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.search.v2
package campaign

import (
    "log"

    "github.com/magicbutton/magic-devices/applogic"
    "github.com/magicbutton/magic-devices/database"
    "github.com/magicbutton/magic-devices/services/models/campaignmodel"
    . "github.com/magicbutton/magic-devices/utils"
)

func CampaignSearch(query string) (*Page[campaignmodel.Campaign], error) {
    log.Println("Calling Campaignsearch")

    return applogic.Search[database.Campaign, campaignmodel.Campaign]("name", query, applogic.MapCampaignOutgoing)

}
