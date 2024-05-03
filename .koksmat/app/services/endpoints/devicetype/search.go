/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.search.v2
package devicetype

import (
    "log"

    "github.com/magicbutton/magic-devices/applogic"
    "github.com/magicbutton/magic-devices/database"
    "github.com/magicbutton/magic-devices/services/models/devicetypemodel"
    . "github.com/magicbutton/magic-devices/utils"
)

func DevicetypeSearch(query string) (*Page[devicetypemodel.Devicetype], error) {
    log.Println("Calling Devicetypesearch")

    return applogic.Search[database.Devicetype, devicetypemodel.Devicetype]("name", query, applogic.MapDevicetypeOutgoing)

}
