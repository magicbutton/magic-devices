/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.search.v2
package importdata

import (
    "log"

    "github.com/magicbutton/magic-devices/applogic"
    "github.com/magicbutton/magic-devices/database"
    "github.com/magicbutton/magic-devices/services/models/importdatamodel"
    . "github.com/magicbutton/magic-devices/utils"
)

func ImportdataSearch(query string) (*Page[importdatamodel.Importdata], error) {
    log.Println("Calling Importdatasearch")

    return applogic.Search[database.Importdata, importdatamodel.Importdata]("name", query, applogic.MapImportdataOutgoing)

}
