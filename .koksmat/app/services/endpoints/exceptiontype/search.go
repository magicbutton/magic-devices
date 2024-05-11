/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.search.v2
package exceptiontype

import (
	"log"

	"github.com/magicbutton/magic-devices/applogic"
	"github.com/magicbutton/magic-devices/database"
	"github.com/magicbutton/magic-devices/services/models/exceptiontypemodel"
	. "github.com/magicbutton/magic-devices/utils"
)

func ExceptiontypeSearch(query string) (*Page[exceptiontypemodel.Exceptiontype], error) {
	log.Println("Calling Exceptiontypesearch")

	return applogic.Search[database.Exceptiontype, exceptiontypemodel.Exceptiontype]("name", query, applogic.MapExceptiontypeOutgoing)

}
