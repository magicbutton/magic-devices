/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/   
//GenerateGoModel v1
package grantedexceptionmodel
import (
	"encoding/json"
	"time"
    // "github.com/magicbutton/magic-devices/database/databasetypes"
)

func UnmarshalGrantedexception(data []byte) (Grantedexception, error) {
	var r Grantedexception
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Grantedexception) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Grantedexception struct {
    ID        int    `json:"id"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
        Tenant string `json:"tenant"`
    Name string `json:"name"`
    Description string `json:"description"`
    Exceptiontype_id int `json:"exceptiontype_id"`
    Person_id int `json:"person_id"`

}

