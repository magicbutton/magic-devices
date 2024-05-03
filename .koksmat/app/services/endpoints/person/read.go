/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: true
---
*/
//generator:  noma3
package person

// noma2
import (
	"log"

	"github.com/magicbutton/magic-devices/applogic"
	"github.com/magicbutton/magic-devices/database"
	"github.com/magicbutton/magic-devices/services/models/applicationmodel"
	"github.com/magicbutton/magic-devices/services/models/personmodel"
)

func PersonRead(id int) (*personmodel.Person, error) {
	log.Println("Calling PersonRead")

	person, err := applogic.Read[database.Person, personmodel.Person](id, applogic.MapPersonOutgoing)
	if err != nil {
		return nil, err
	}
	applications, err := applogic.Select[database.Application, applicationmodel.Application]("owner_id = ?", applogic.MapApplicationOutgoing, person.ID)
	if err != nil {
		return nil, err
	}
	person.Applications = *applications
	return person, nil

}
