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
	"io/ioutil"
	"log"
	"net/http"
	// "net/url"
	"encoding/json"
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
		"collections": map[string]map[string]int{
			"templates":     map[string]int{},
			"surveys":       map[string]int{"collectors": 1, "invites": 1, "responses": 1, "structure": 1, "invite_codes": 1, "groups": 1, "reports": 1, "csv": 1},
			"collectors":    map[string]int{},
			"contacts":      map[string]int{},
			"embed":         map[string]int{},
			"contact-lists": map[string]int{"contacts": 1},
			"webhooks":      map[string]int{},
		},
	}
	FORMAT = "" // default is json
)

func Create(entityType string, data map[string]string) string {
	client := &http.Client{}
	URL := HOST + entityType + "/"

	b, err := json.Marshal(data)

	//pass the values to the request's body
	req, err := http.NewRequest("POST", URL, bytes.NewReader(b))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	req.SetBasicAuth(EMAIL, PASSWORD)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	bodyText, err := ioutil.ReadAll(resp.Body)
	s := string(bodyText)
	return s
}

func Get(entityType, id string) string {
	client := &http.Client{}
	URL := HOST + entityType + "/" + id + "/"

	req, err := http.NewRequest("GET", URL, nil)
	req.SetBasicAuth(EMAIL, PASSWORD)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	s := string(bodyText)
	return s
}

func Update(entityType, id string, data map[string]string) string {
	client := &http.Client{}
	URL := HOST + entityType + "/" + id + "/"

	b, err := json.Marshal(data)

	//pass the values to the request's body
	req, err := http.NewRequest("PUT", URL, bytes.NewReader(b))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	req.SetBasicAuth(EMAIL, PASSWORD)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	bodyText, err := ioutil.ReadAll(resp.Body)
	s := string(bodyText)
	return s
}

func Delete(entityType, id string) string {
	client := &http.Client{}
	URL := HOST + entityType + "/" + id + "/"

	req, err := http.NewRequest("DELETE", URL, nil)
	req.SetBasicAuth(EMAIL, PASSWORD)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	s := string(bodyText)
	return s
}

func List(entityType string, args map[string]string) string {
	client := &http.Client{}
	URL := HOST + entityType + "/"

	if args != nil {
		URL = URL + "?"
	}
	for key, value := range args {
		URL = URL + key + "=" + value
	}
	// IF args are not nil then add to query string plus add predicate, etc.

	req, err := http.NewRequest("GET", URL, nil)
	req.SetBasicAuth(EMAIL, PASSWORD)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	s := string(bodyText)
	return s
}

func main() {

	log.Println("Running..")
	// log.Println(ListResponses("550996"))

	// log.Println(FilterCSVByDateUpdated("550996", "2013-11-04", ">"))
}
