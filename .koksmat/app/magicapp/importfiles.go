package magicapp

import (
	"fmt"
	"os"

	"github.com/magicbutton/magic-devices/services/endpoints/importdata"
	"github.com/magicbutton/magic-devices/services/models/importdatamodel"
)

func ImportFiles(filepath string) (*importdatamodel.Importdata, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {

		return nil, fmt.Errorf("could not read file: %w", err)
	}

	OpenDatabase()
	text := string(data)

	importRecord := importdatamodel.Importdata{
		Name:        filepath,
		Description: "Imported file",
		User_id:     1,
		Tenant:      "",
		Data:        []byte(text),
	}
	return importdata.ImportdataCreate(importRecord)

}
