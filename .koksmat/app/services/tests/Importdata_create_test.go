    /* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/
    //generator:  noma4.1
    package tests
    import (
        "testing"
        "github.com/magicbutton/magic-devices/services/endpoints/importdata"
                    "github.com/magicbutton/magic-devices/services/models/importdatamodel"
        "github.com/stretchr/testify/assert"
    )
    
    func TestImportdatacreate(t *testing.T) {
                                // transformer v1
            object := importdatamodel.Importdata{}
         
            result,err := importdata.ImportdataCreate(object)
            if err != nil {
                t.Errorf("Error %s", err)
            }
            assert.NotNil(t, result)
        
    
    }
    
