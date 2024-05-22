/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.search.v2
package person

import (
	"log"

	"github.com/magicbutton/magic-devices/applogic"
	"github.com/magicbutton/magic-devices/database"
	"github.com/magicbutton/magic-devices/services/models/personmodel"
	. "github.com/magicbutton/magic-devices/utils"
)

func PersonSearch(query string) (*Page[personmodel.Person], error) {
	log.Println("Calling Personsearch")

	return applogic.Search[database.Person, personmodel.Person]("name", query, applogic.MapPersonOutgoing)

}
