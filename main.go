/**
 * fluidsurveys
 *
 * FluidSurveys API wrapper
 *
 * @author Collin Glass <collin@fluidware.com>
 * @copyright 2014 Fluidware
 * @license MIT <https://raw.github.com/fluidware/fluidsurveys-node/master/LICENSE>
 * @link http://fluidsurveys.com
 * @version 0.1.3
 */

package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
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

func NewSurvey(name string) string {
	client := &http.Client{}
	URL := HOST + "surveys/"
	if FORMAT != "" {
		URL = URL + "&format=" + FORMAT
	}

	v := url.Values{}
	v.Set("name", name)
	//pass the values to the request's body
	req, err := http.NewRequest("POST", URL, strings.NewReader(v.Encode()))
	if err != nil {
		log.Fatal(err)
	}

	req.SetBasicAuth(EMAIL, PASSWORD)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	s := string(bodyText)
	return s
}

func List() string {
	client := &http.Client{}
	url := HOST + "surveys/"
	if FORMAT != "" {
		url = url + "?format=" + FORMAT
	}
	req, err := http.NewRequest("GET", url, nil)
	req.SetBasicAuth(EMAIL, PASSWORD)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	s := string(bodyText)
	return s
}

func ListResponses(id string) string {
	client := &http.Client{}
	url := HOST + "surveys/" + id + "/repsonses/"
	if FORMAT != "" {
		url = url + "?format=" + FORMAT
	}
	req, err := http.NewRequest("GET", url, nil)
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

	log.Println("Initial Commit.")
	log.Println(NewSurvey("TESTING"))
}
