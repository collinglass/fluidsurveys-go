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

package fluidsurveys

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
	EMAIL    = "" // your email/username
	PASSWORD = "" // your Password

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

//	Fluidsurveys(...) takes your username and password and sets it package wide
//	username: Your username
//	password: Your password
func Fluidsurveys(username, password string) {
	EMAIL = username
	PASSWORD = password
}

//	SetHost(...) takes your host name and sets it package wide
//	host: Your host api url
func SetHost(host string) {
	HOST = host
}

// Create(...) creates a new entity
// entityType: Entity type you want to create
// data: map of data to set on new entity
func Create(entityType string, data map[string]interface{}) ([]byte, error) {
	resp, err := makeRequest("POST", fmt.Sprintf("%s%s/", HOST, checkEntity(entityType)), data)

	return resp, err
}

// CreateChild(...) creates a new child entity on a collection
// parentType: Parent type
// parentId: Parent Id
// childType: Type of child you want to create
// childId: Id of child you want to create
// data: map of data you want to update on entity
func CreateChild(parentType string, parentId uint64, childType string, data map[string]interface{}) ([]byte, error) {
	resp, err := makeRequest("POST", fmt.Sprintf("%s%s/%d/%s/", HOST, checkEntity(parentType), parentId, checkChild(parentType, childType)), data)

	return resp, err
}

// Delete(...) deletes an entity
// entityType: Entity type you want to delete
// id: Id of entity you want to delete
func Delete(entityType string, id uint64) ([]byte, error) {
	resp, err := makeRequest("DELETE", fmt.Sprintf("%s%s/%d/", HOST, checkEntity(entityType), id), nil)

	return resp, err
}

// Get(...) gets an entity by id
// entityType: Entity type you want to get
// id: Id of entity you want to get
func Get(entityType string, id uint64) ([]byte, error) {
	var url string
	if entityType == "contact-lists" {
		url = fmt.Sprintf("%s%s/%d/contacts/", HOST, checkEntity(entityType), id)
	} else {
		url = fmt.Sprintf("%s%s/%d/", HOST, checkEntity(entityType), id)
	}
	resp, err := makeRequest("GET", url, nil)

	return resp, err
}

// GetChild(...) gets a child entity by id
// parentType: Parent type
// parentId: Parent Id
// childType: Type of child you want to get
// childId: Id of child you want to get
func GetChild(parentType string, parentId uint64, childType string, childId uint64) ([]byte, error) {
	resp, err := makeRequest("GET", fmt.Sprintf("%s%s/%d/%s/%d/", HOST, checkEntity(parentType), parentId, checkChild(parentType, childType), childId), nil)

	return resp, err
}

// List(...) gets a list of entities
// entityType: Entity type you want to get
// args: arguments to filter results
func List(entityType string, args map[string]string) ([]byte, error) {
	URL, err := url.Parse(fmt.Sprintf("%s%s/", HOST, checkEntity(entityType)))
	if err != nil {
		return nil, err
	}

	v := URL.Query()
	for key, value := range args {
		v.Set(key, value)
	}
	URL.RawQuery = v.Encode()

	resp, err := makeRequest("GET", URL.String(), nil)

	return resp, err
}

// ListChild(...) gets a list of children
// parentType: Parent type
// parentId: Parent Id
// childType: Type of children you want to get
func ListChildren(parentType string, parentId uint64, childType string, args map[string]string) ([]byte, error) {
	URL, err := url.Parse(fmt.Sprintf("%s%s/%d/%s/", HOST, checkEntity(parentType), parentId, checkChild(parentType, childType)))
	if err != nil {
		return nil, err
	}

	v := URL.Query()
	for key, value := range args {
		v.Set(key, value)
	}
	URL.RawQuery = v.Encode()

	resp, err := makeRequest("GET", URL.String(), nil)

	return resp, err
}

// Update(...) updates an entity
// entityType: Entity type you want to update
// id: Id of entity you want to update
// data: map of data you want to update on entity
func Update(entityType string, id uint64, data map[string]interface{}) ([]byte, error) {
	if entityType == "surveys" {
		var result []byte
		var err error
		if data["name"] != nil {
			result, err = makeRequest("PUT", fmt.Sprintf("%s%s/%d/", HOST, checkEntity(entityType), id), data)
		}
		if data["structure"] != nil {
			result, err = makeRequest("PUT", fmt.Sprintf("%s%s/%d/structure/", HOST, checkEntity(entityType), id), data)
		}
		return result, err
	}
	resp, err := makeRequest("PUT", fmt.Sprintf("%s%s/%d/", HOST, checkEntity(entityType), id), data)

	return resp, err
}

// makeRequest(...) makes all requests
// method: GET, POST, PUT, DELETE, etc.
// urlString: url to send request
// data: map of data to use with request
func makeRequest(method, urlString string, data map[string]interface{}) ([]byte, error) {
	//pass the values to the request's body
	var bodyReader *bytes.Reader
	if data != nil {
		body, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		bodyReader = bytes.NewReader(body)
	} else {
		bodyReader = bytes.NewReader(nil)
	}

	req, err := http.NewRequest(method, urlString, bodyReader)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(EMAIL, PASSWORD)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return bodyText, err
}

// checkEntity(...) checks if entity is a valid collection
func checkEntity(entityType string) string {
	if collections[entityType] == nil {
		log.Fatal(errors.New("Invalid Entity Type"))
	}
	return entityType
}

// checkChild(...) checks if child is a valid child of a collection
func checkChild(entityType, childType string) string {
	if collections[entityType][childType] != 1 {
		log.Fatal(errors.New("Invalid Child Type"))
	}
	return childType
}
