package fluidsurveys

import ()

func main() {

	log.Println("Running..")

	Setup("", "")
	// data := map[string]interface{}{"name": "New Survey123", "structure": map[string]interface{}{}}
	data := map[string]interface{}{
		"name": "This is",
		"structure": map[string]interface{}{
			"id": 554106, "created_at": "2014-06-18T23:39:23Z",
			"deploy_url":           "http://fluidsurveys.com/surveys/collin-n2c/new-survey123-13/",
			"survey_structure_uri": "https://fluidsurveys.com/api/v3/surveys/554106/structure/",
			"collectors_uri":       "https://fluidsurveys.com/api/v3/surveys/554106/collectors/",
			"number_of_responses":  0, "creator": "https://fluidsurveys.com/api/v3/users/2127403118/",
			"send_invite_uri":  "https://fluidsurveys.com/api/v3/surveys/554106/invites/",
			"responses_uri":    "https://fluidsurveys.com/api/v3/surveys/554106/responses/",
			"invite_codes_uri": "https://fluidsurveys.com/api/v3/surveys/554106/invite_codes/",
			"reports_uri":      "https://fluidsurveys.com/api/v3/surveys/554106/reports/",
			"csv_uri":          "https://fluidsurveys.com/api/v3/surveys/554106/csv/",
			"survey_uri":       "https://fluidsurveys.com/api/v3/surveys/554106/",
			"live":             1, "name": "New Survey123",
			"slug":       "new-survey123-13",
			"updated_at": "2014-06-18T23:39:23Z",
			"groups_uri": "https://fluidsurveys.com/api/v3/surveys/554106/groups/",
		},
	}

	log.Println(Update("surveys", 554134, data))
	// log.Println(Get("surveys", 554106))
}
