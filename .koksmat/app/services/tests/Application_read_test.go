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
        
        "github.com/stretchr/testify/assert"
    )
    
    func TestApplicationread(t *testing.T) {
                    
            result,err := application.ApplicationRead(-1)
            if err != nil {
                t.Errorf("Error %s", err)
            }
            assert.NotNil(t, result)
        
    
    }
    
