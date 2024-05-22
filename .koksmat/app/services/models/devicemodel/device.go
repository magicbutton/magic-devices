/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/
//GenerateGoModel v1
package devicemodel

import (
	"encoding/json"
	"time"
)

func UnmarshalDevice(data []byte) (Device, error) {
	var r Device
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Device) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Device struct {
	ID                          int       `json:"id"`
	CreatedAt                   time.Time `json:"created_at"`
	UpdatedAt                   time.Time `json:"updated_at"`
	Tenant                      string    `json:"tenant"`
	Name                        string    `json:"name"`
	Description                 string    `json:"description"`
	Unique_Device_Id            string    `json:"unique_device_id"`
	Displayname                 string    `json:"displayname"`
	Macaddress                  string    `json:"macaddress"`
	Devicetype_id               int       `json:"devicetype_id"`
	Antivirus                   bool      `json:"antivirus"`
	Dlp                         bool      `json:"dlp"`
	Endpointdetection           bool      `json:"endpointdetection"`
	Diskentryptionprimarydisk   bool      `json:"diskentryptionprimarydisk"`
	Diskentryptionsecondarydisk bool      `json:"diskentryptionsecondarydisk"`
}
