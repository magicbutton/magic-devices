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
	"github.com/magicbutton/magic-devices/services/models/campaignmodel"
   
)


func MapCampaignOutgoing(db database.Campaign) campaignmodel.Campaign {
    return campaignmodel.Campaign{
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

    }
}

func MapCampaignIncoming(in campaignmodel.Campaign) database.Campaign {
    return database.Campaign{
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

    }
}
