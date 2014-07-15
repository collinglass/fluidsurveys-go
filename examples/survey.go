package main

import (
	"encoding/json"
	"github.com/collinglass/fluidsurveys-go/fluidsurveys"
	"log"
)

func main() {

	log.Println("Running..")

	fluidsurveys.Fluidsurveys("your_email", "your_password")

	data := map[string]interface{}{
		"name": "Demo Survey 2",
		"structure": map[string]interface{}{
			"created_at":           "2014-06-18T23:39:23Z",
			"survey_structure_uri": "https://fluidsurveys.com/api/v3/surveys/123456/structure/",
			"collectors_uri":       "https://fluidsurveys.com/api/v3/surveys/123456/collectors/",
			"number_of_responses":  0,
			"send_invite_uri":      "https://fluidsurveys.com/api/v3/surveys/123456/invites/",
			"responses_uri":        "https://fluidsurveys.com/api/v3/surveys/123456/responses/",
			"invite_codes_uri":     "https://fluidsurveys.com/api/v3/surveys/123456/invite_codes/",
			"reports_uri":          "https://fluidsurveys.com/api/v3/surveys/123456/reports/",
			"csv_uri":              "https://fluidsurveys.com/api/v3/surveys/123456/csv/",
			"survey_uri":           "https://fluidsurveys.com/api/v3/surveys/123456/",
			"live":                 1,
			"name":                 "New Survey123",
			"slug":                 "new-survey123-13",
			"updated_at":           "2014-06-18T23:39:23Z",
			"groups_uri":           "https://fluidsurveys.com/api/v3/surveys/123456/groups/",
		},
	}

	demoSurvey(data)
}

func demoSurvey(data map[string]interface{}) {
	// Create survey
	result, err := fluidsurveys.Create("surveys", data)
	if err != nil {
		log.Println(err)
	}

	// New survey object for unmarshalling
	var new_survey map[string]interface{}

	// Unmarshall response into new_survey
	err = json.Unmarshal(result, &new_survey)
	if err != nil {
		log.Println(err)
	}

	// Take id from new survey and convert to uint64
	idf, _ := new_survey["id"].(float64)
	id := uint64(idf)

	// Get survey
	result, err = fluidsurveys.Get("surveys", id)

	// Survey object
	var survey map[string]interface{}

	// Unmarshall response into update object
	err = json.Unmarshal(result, &survey)
	if err != nil {
		log.Println(err)
	}

	log.Println(survey)

	// Update survey
	result, err = fluidsurveys.Update("surveys", id, data)
	if err != nil {
		log.Println(err)
	}
	// Update message
	var updated string

	// Unmarshal response into update object
	err = json.Unmarshal(result, &updated)
	if err != nil {
		log.Println(err)
	}

	log.Println(updated)

	// Delete survey
	result, err = fluidsurveys.Delete("surveys", id)
	if err != nil {
		log.Println(err)
	}
	log.Println(result)

	// Deleted message
	var deleted string

	// Unmarshal response into deleted message
	err = json.Unmarshal(result, &deleted)
	if err != nil {
		log.Println(err)
	}

	log.Println(deleted)
}
