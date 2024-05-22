/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.search.v2
package grantedexception

import (
    "log"

    "github.com/magicbutton/magic-devices/applogic"
    "github.com/magicbutton/magic-devices/database"
    "github.com/magicbutton/magic-devices/services/models/grantedexceptionmodel"
    . "github.com/magicbutton/magic-devices/utils"
)

func GrantedexceptionSearch(query string) (*Page[grantedexceptionmodel.Grantedexception], error) {
    log.Println("Calling Grantedexceptionsearch")

    return applogic.Search[database.Grantedexception, grantedexceptionmodel.Grantedexception]("name", query, applogic.MapGrantedexceptionOutgoing)

}
