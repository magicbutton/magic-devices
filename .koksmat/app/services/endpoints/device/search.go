/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.search.v2
package device

import (
	"log"

	"github.com/magicbutton/magic-devices/applogic"
	"github.com/magicbutton/magic-devices/database"
	"github.com/magicbutton/magic-devices/services/models/devicemodel"
	. "github.com/magicbutton/magic-devices/utils"
)

func DeviceSearch(query string) (*Page[devicemodel.Device], error) {
	log.Println("Calling Devicesearch")

	return applogic.Search[database.Device, devicemodel.Device]("name", query, applogic.MapDeviceOutgoing)

}
