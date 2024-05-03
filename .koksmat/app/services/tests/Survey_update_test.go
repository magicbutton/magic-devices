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
        "github.com/magicbutton/magic-devices/services/endpoints/survey"
                    "github.com/magicbutton/magic-devices/services/models/surveymodel"
        "github.com/stretchr/testify/assert"
    )
    
    func TestSurveyupdate(t *testing.T) {
                                // transformer v1
            object := surveymodel.Survey{}
         
            result,err := survey.SurveyUpdate(object)
            if err != nil {
                t.Errorf("Error %s", err)
            }
            assert.NotNil(t, result)
        
    
    }
    
