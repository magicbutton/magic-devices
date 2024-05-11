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
	"github.com/magicbutton/magic-devices/services/models/grantedexceptionmodel"
)

func MapGrantedexceptionOutgoing(db database.Grantedexception) grantedexceptionmodel.Grantedexception {
	return grantedexceptionmodel.Grantedexception{
		ID:               db.ID,
		CreatedAt:        db.CreatedAt,
		UpdatedAt:        db.UpdatedAt,
		Tenant:           db.Tenant,
		Name:             db.Name,
		Description:      db.Description,
		Exceptiontype_id: db.Exceptiontype_id,
		Person_id:        db.Person_id,
	}
}

func MapGrantedexceptionIncoming(in grantedexceptionmodel.Grantedexception) database.Grantedexception {
	return database.Grantedexception{
		ID:               in.ID,
		CreatedAt:        in.CreatedAt,
		UpdatedAt:        in.UpdatedAt,
		Tenant:           in.Tenant,
		Name:             in.Name,
		Description:      in.Description,
		Exceptiontype_id: in.Exceptiontype_id,
		Person_id:        in.Person_id,
	}
}
