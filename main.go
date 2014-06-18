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
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

const (
	HOST     = "https://fluidsurveys.com/api/v3/"
	EMAIL    = "" // Enter your email/username here
	PASSWORD = "" // Enter your Password here
)

var (
	routes = map[string]interface{}{
		"types": map[string]string{
			"root":    "/{type}/",
			"details": "/{type}/{id:[0-9]+}/",
		},
		"collections": map[string][]string{
			"templates":     []string{},
			"surveys":       []string{"collectors", "invites", "responses", "structure", "invite_codes", "groups", "reports", "csv"},
			"collectors":    []string{},
			"contacts":      []string{},
			"embed":         []string{},
			"contact-lists": []string{"contacts"},
			"webhooks":      []string{},
		},
	}
	FORMAT = "" // default is json
)

func makeRequest(method, urlString string, data map[string]string) (string, error) {
	client := &http.Client{}

	//pass the values to the request's body
	var bodyReader *bytes.Reader
	if data != nil {
		body, err := json.Marshal(data)
		handle(err)
		bodyReader = bytes.NewReader(body)
	} else {
		bodyReader = bytes.NewReader([]byte{})
	}

	log.Println(urlString)

	req, err := http.NewRequest(method, urlString, bodyReader)
	handle(err)

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(EMAIL, PASSWORD)

	resp, err := client.Do(req)
	handle(err)

	bodyText, err := ioutil.ReadAll(resp.Body)
	handle(err)

	s := string(bodyText)

	return s, err
}

func handle(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func Create(entityType string, data map[string]string) (string, error) {
	return makeRequest("POST", HOST+entityType, data)
}

func Get(entityType, id string) (string, error) {
	return makeRequest("GET", HOST+entityType+"/"+id+"/", nil)
}

func Update(entityType, id string, data map[string]string) (string, error) {
	return makeRequest("PUT", HOST+entityType+"/"+id+"/", data)
}

func Delete(entityType, id string) (string, error) {
	return makeRequest("DELETE", HOST+entityType+"/"+id+"/", nil)
}

func List(entityType string, args map[string]string) (string, error) {
	URL, err := url.Parse(HOST + "/" + entityType + "/")
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

	log.Println(List("surveys", nil))
}
