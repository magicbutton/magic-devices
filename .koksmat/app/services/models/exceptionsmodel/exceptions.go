/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/
//GenerateGoModel v1
package exceptionsmodel

import (
	"encoding/json"
	"time"
)

func UnmarshalException(data []byte) (Exception, error) {
	var r Exception
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Exception) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Exception struct {
	ID               int       `json:"id"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	Tenant           string    `json:"tenant"`
	Name             string    `json:"name"`
	Description      string    `json:"description"`
	Exceptiontype_id int       `json:"exceptiontype_id"`
	Person_id        int       `json:"person_id"`
}
