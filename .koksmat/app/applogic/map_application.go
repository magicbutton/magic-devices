/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: true
---
*/
//GenerateMapModel v1.1
package applogic

import (
	//"encoding/json"
	//"time"
	"github.com/magicbutton/magic-devices/database"
	"github.com/magicbutton/magic-devices/services/models/applicationmodel"
)

func MapApplicationOutgoing(item database.Application) applicationmodel.Application {

	return applicationmodel.Application{
		ID:          item.ID,
		CreatedAt:   item.CreatedAt,
		UpdatedAt:   item.UpdatedAt,
		Name:        item.Name,
		Description: item.Description,
		Key:         item.Key,
		Displayname: item.Displayname,
		Owner_id:    item.Owner_id,
	}
}

func MapApplicationIncoming(item applicationmodel.Application) database.Application {
	return database.Application{
		ID:          item.ID,
		CreatedAt:   item.CreatedAt,
		UpdatedAt:   item.UpdatedAt,
		Name:        "displayname: " + item.Displayname + " key: " + item.Key,
		Description: item.Description,
		Key:         item.Key,
		Displayname: item.Displayname,
		Owner_id:    item.Owner_id,
	}
}
