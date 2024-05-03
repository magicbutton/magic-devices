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
	"github.com/magicbutton/magic-devices/services/models/devicemodel"
   
)


func MapDeviceOutgoing(db database.Device) devicemodel.Device {
    return devicemodel.Device{
        ID:        db.ID,
        CreatedAt: db.CreatedAt,
        UpdatedAt: db.UpdatedAt,
                Tenant : db.Tenant,
        Name : db.Name,
        Description : db.Description,
        Unique_Device_Id : db.Unique_Device_Id,
        Displayname : db.Displayname,
        Macaddress : db.Macaddress,
                Devicetype_id : db.Devicetype_id,

    }
}

func MapDeviceIncoming(in devicemodel.Device) database.Device {
    return database.Device{
        ID:        in.ID,
        CreatedAt: in.CreatedAt,
        UpdatedAt: in.UpdatedAt,
                Tenant : in.Tenant,
        Name : in.Name,
        Description : in.Description,
        Unique_Device_Id : in.Unique_Device_Id,
        Displayname : in.Displayname,
        Macaddress : in.Macaddress,
                Devicetype_id : in.Devicetype_id,

    }
}
