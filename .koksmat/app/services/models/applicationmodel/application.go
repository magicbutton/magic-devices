/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/   
//GenerateGoModel v1
package applicationmodel
import (
	"encoding/json"
	"time"
    // "github.com/magicbutton/magic-devices/database/databasetypes"
)

func UnmarshalApplication(data []byte) (Application, error) {
	var r Application
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Application) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Application struct {
    ID        int    `json:"id"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
        Tenant string `json:"tenant"`
    Name string `json:"name"`
    Description string `json:"description"`
    Key string `json:"key"`
    Displayname string `json:"displayname"`
    Imported_Ownername string `json:"imported_ownername"`
    Owner_id int `json:"owner_id"`

}

