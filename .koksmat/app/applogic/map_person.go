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
	"github.com/magicbutton/magic-devices/database"
	"github.com/magicbutton/magic-devices/services/models/personmodel"
)

func MapPersonOutgoing(item database.Person) personmodel.Person {
	return personmodel.Person{
		ID:        item.ID,
		CreatedAt: item.CreatedAt,
		UpdatedAt: item.UpdatedAt,

		Name:             item.Name,
		Unique_Person_Id: item.Unique_Person_Id,
		Email:            item.Email,
		Displayname:      item.Displayname,
		Description:      item.Description,
	}
}

func MapPersonIncoming(item personmodel.Person) database.Person {

	return database.Person{
		ID:               item.ID,
		CreatedAt:        item.CreatedAt,
		UpdatedAt:        item.UpdatedAt,
		Displayname:      item.Displayname,
		Name:             " email: " + item.Email + " uniqueid: " + item.Unique_Person_Id + " displayname: " + item.Displayname,
		Description:      item.Description,
		Unique_Person_Id: item.Unique_Person_Id,

		Email: item.Email,
	}
}
