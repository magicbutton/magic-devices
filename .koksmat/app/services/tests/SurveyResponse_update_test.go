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
        "github.com/magicbutton/magic-devices/services/endpoints/surveyresponse"
                    "github.com/magicbutton/magic-devices/services/models/surveyresponsemodel"
        "github.com/stretchr/testify/assert"
    )
    
    func TestSurveyResponseupdate(t *testing.T) {
                                // transformer v1
            object := surveyresponsemodel.SurveyResponse{}
         
            result,err := surveyresponse.SurveyResponseUpdate(object)
            if err != nil {
                t.Errorf("Error %s", err)
            }
            assert.NotNil(t, result)
        
    
    }
    
