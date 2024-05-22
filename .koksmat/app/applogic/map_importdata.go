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
	"github.com/magicbutton/magic-devices/database"
	"github.com/magicbutton/magic-devices/services/models/importdatamodel"
)

func MapImportdataOutgoing(db database.Importdata) importdatamodel.Importdata {
	return importdatamodel.Importdata{
		ID:          db.ID,
		CreatedAt:   db.CreatedAt,
		UpdatedAt:   db.UpdatedAt,
		Tenant:      db.Tenant,
		Name:        db.Name,
		Description: db.Description,
		User_id:     db.User_id,
	}
}

func MapImportdataIncoming(in importdatamodel.Importdata) database.Importdata {
	return database.Importdata{
		ID:          in.ID,
		CreatedAt:   in.CreatedAt,
		UpdatedAt:   in.UpdatedAt,
		Tenant:      in.Tenant,
		Name:        in.Name,
		Description: in.Description,
		User_id:     in.User_id,
	}
}
