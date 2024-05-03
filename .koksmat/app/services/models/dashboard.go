package models

type Dashboard struct {
	PersonId                int      `json:"person_id"`
	NumberOfApplications    int      `json:"number_of_applications"`
	NumberOfSurveys         int      `json:"number_of_surveys"`
	NumberOfSurveyResponses int      `json:"number_of_survey_responses"`
	NumberOfAppsYouOwn      int      `json:"number_of_apps_you_own"`
	NumberOfOwners          int      `json:"number_of_owners"`
	Messages                []string `json:"messages"`
}
