/**
 * fluidsurveys
 *
 * FluidSurveys API wrapper
 *
 * @author Collin Glass <collin@fluidware.com>
 * @copyright 2014 Fluidware
 * @license MIT <https://raw.github.com/collinglass/fluidsurveys-go/master/LICENSE>
 * @link http://fluidsurveys.com
 * @version 0.1.3
 */

package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

var (
	HOST     = "https://fluidsurveys.com/api/v3/"
	EMAIL    = "" // Enter your email/username here
	PASSWORD = "" // Enter your Password here

	collections = map[string]map[string]uint8{
		"templates":     map[string]uint8{},
		"surveys":       map[string]uint8{"collectors": 1, "invites": 1, "responses": 1, "structure": 1, "invite_codes": 1, "groups": 1, "reports": 1, "csv": 1},
		"collectors":    map[string]uint8{},
		"contacts":      map[string]uint8{},
		"embed":         map[string]uint8{},
		"contact-lists": map[string]uint8{"contacts": 1},
		"webhooks":      map[string]uint8{},
	}
	FORMAT = "" // default is json
)

func Setup(username, password string) {
	EMAIL = username
	PASSWORD = password
}

func makeRequest(method, urlString string, data map[string]interface{}) (map[string]interface{}, error) {

	//pass the values to the request's body
	var bodyReader *bytes.Reader
	if data != nil {
		body, err := json.Marshal(data)
		handle(err)
		bodyReader = bytes.NewReader(body)
	} else {
		bodyReader = bytes.NewReader(nil)
	}

	req, err := http.NewRequest(method, urlString, bodyReader)
	handle(err)

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(EMAIL, PASSWORD)

	resp, err := http.DefaultClient.Do(req)
	handle(err)

	bodyText, err := ioutil.ReadAll(resp.Body)
	handle(err)

	result := map[string]interface{}{}

	err = json.Unmarshal(bodyText, &result)
	if err != nil {
		result["response"] = string(bodyText)
		err = nil
	}

	return result, err
}

func handle(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func checkEntity(entityType string) string {
	if collections[entityType] == nil {
		log.Fatal(errors.New("Invalid Entity Type"))
	}
	return entityType
}

func checkChild(entityType, childType string) string {
	if collections[entityType][childType] != 1 {
		log.Fatal(errors.New("Invalid Child Type"))
	}
	return childType
}

func Create(entityType string, data map[string]interface{}) (map[string]interface{}, error) {
	if entityType == "surveys" {
		resp, err := makeRequest("POST", fmt.Sprintf("%s%s/", HOST, checkEntity(entityType)), data)
		handle(err)

		idf, _ := resp["id"].(float64)
		id := int(idf)
		return makeRequest("PUT", fmt.Sprintf("%s%s/%d/structure/", HOST, checkEntity(entityType), id), data)
	}
	return makeRequest("POST", fmt.Sprintf("%s%s/", HOST, checkEntity(entityType)), data)
}

func Get(entityType string, id uint64) (map[string]interface{}, error) {
	return makeRequest("GET", fmt.Sprintf("%s%s/%d/", HOST, checkEntity(entityType), id), nil)
}

func Update(entityType string, id uint64, data map[string]interface{}) (map[string]interface{}, error) {
	if entityType == "surveys" {
		var result map[string]interface{}
		var err error
		if data["name"] != nil {
			result, err = makeRequest("PUT", fmt.Sprintf("%s%s/%d/", HOST, checkEntity(entityType), id), data)
		}
		if data["structure"] != nil {
			result, err = makeRequest("PUT", fmt.Sprintf("%s%s/%d/structure/", HOST, checkEntity(entityType), id), data)
		}
		return result, err
	}
	return makeRequest("PUT", fmt.Sprintf("%s%s/%d/", HOST, checkEntity(entityType), id), data)
}

func Delete(entityType string, id uint64) (map[string]interface{}, error) {
	return makeRequest("DELETE", fmt.Sprintf("%s%s/%d/", HOST, checkEntity(entityType), id), nil)
}

func List(entityType string, args map[string]string) (map[string]interface{}, error) {
	URL, err := url.Parse(fmt.Sprintf("%s%s/", HOST, checkEntity(entityType)))
	handle(err)

	v := URL.Query()
	for key, value := range args {
		v.Set(key, value)
	}
	URL.RawQuery = v.Encode()

	return makeRequest("GET", URL.String(), nil)
}

func CreateChild(parentType string, parentId uint64, childType string, data map[string]interface{}) (map[string]interface{}, error) {
	return makeRequest("POST", fmt.Sprintf("%s%s/%d/%s/", HOST, checkEntity(parentType), parentId, checkChild(parentType, childType)), data)
}

func GetChild(parentType string, parentId uint64, childType string, childId uint64) (map[string]interface{}, error) {
	return makeRequest("GET", fmt.Sprintf("%s%s/%d/%s/%d/", HOST, checkEntity(parentType), parentId, checkChild(parentType, childType), childId), nil)
}

func ListChildren(parentType string, parentId uint64, childType string, childId uint64, args map[string]string) (map[string]interface{}, error) {
	URL, err := url.Parse(fmt.Sprintf("%s%s/%d/%s/%d/", HOST, checkEntity(parentType), parentId, checkChild(parentType, childType), childId))
	handle(err)

	v := URL.Query()
	for key, value := range args {
		v.Set(key, value)
	}
	URL.RawQuery = v.Encode()

	return makeRequest("GET", URL.String(), nil)
}

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
