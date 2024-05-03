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
	//"encoding/json"
	//"time"
	"github.com/magicbutton/magic-devices/database"
	"github.com/magicbutton/magic-devices/services/models/exceptiontypemodel"
   
)


func MapExceptiontypeOutgoing(db database.Exceptiontype) exceptiontypemodel.Exceptiontype {
    return exceptiontypemodel.Exceptiontype{
        ID:        db.ID,
        CreatedAt: db.CreatedAt,
        UpdatedAt: db.UpdatedAt,
                Tenant : db.Tenant,
        Name : db.Name,
        Description : db.Description,
        Family : db.Family,

    }
}

func MapExceptiontypeIncoming(in exceptiontypemodel.Exceptiontype) database.Exceptiontype {
    return database.Exceptiontype{
        ID:        in.ID,
        CreatedAt: in.CreatedAt,
        UpdatedAt: in.UpdatedAt,
                Tenant : in.Tenant,
        Name : in.Name,
        Description : in.Description,
        Family : in.Family,

    }
}
