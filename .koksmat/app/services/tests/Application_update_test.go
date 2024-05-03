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
        "github.com/magicbutton/magic-devices/services/endpoints/application"
                    "github.com/magicbutton/magic-devices/services/models/applicationmodel"
        "github.com/stretchr/testify/assert"
    )
    
    func TestApplicationupdate(t *testing.T) {
                                // transformer v1
            object := applicationmodel.Application{}
         
            result,err := application.ApplicationUpdate(object)
            if err != nil {
                t.Errorf("Error %s", err)
            }
            assert.NotNil(t, result)
        
    
    }
    
