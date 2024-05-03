/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/   
//GenerateGoModel v1
package campaignmodel
import (
	"encoding/json"
	"time"
    // "github.com/magicbutton/magic-devices/database/databasetypes"
)

func UnmarshalCampaign(data []byte) (Campaign, error) {
	var r Campaign
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Campaign) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Campaign struct {
    ID        int    `json:"id"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
        Tenant string `json:"tenant"`
    Name string `json:"name"`
    Description string `json:"description"`
    Url string `json:"url"`
    Key string `json:"key"`
    Displayname string `json:"displayname"`
    Owner_id int `json:"owner_id"`

}

