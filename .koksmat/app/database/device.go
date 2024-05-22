/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/
//version: pølsevogn2
package database

import (
	"time"

	"github.com/uptrace/bun"
)

type Device struct {
	bun.BaseModel `bun:"table:device,alias:device"`

	ID                          int       `bun:"id,pk,autoincrement"`
	CreatedAt                   time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt                   time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	DeletedAt                   time.Time `bun:",soft_delete,nullzero"`
	Tenant                      string    `bun:"tenant"`
	Name                        string    `bun:"name"`
	Description                 string    `bun:"description"`
	Unique_Device_Id            string    `bun:"unique_device_id"`
	Displayname                 string    `bun:"displayname"`
	Macaddress                  string    `bun:"macaddress"`
	Devicetype_id               int       `bun:"devicetype_id"`
	Antivirus                   bool      `bun:"antivirus"`
	Dlp                         bool      `bun:"dlp"`
	Endpointdetection           bool      `bun:"endpointdetection"`
	Diskentryptionprimarydisk   bool      `bun:"diskentryptionprimarydisk"`
	Diskentryptionsecondarydisk bool      `bun:"diskentryptionsecondarydisk"`
}
