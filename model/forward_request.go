package model

import (
"log"
"io/ioutil"
"bytes"
"fmt"
"net/http"
"net/url"
"encoding/json"
"os"
)

type APIResponse struct {
	Body interface{} `json:"response"`
}

type ForwardRequest struct {
	Action  string  `json:"action"`
	Payload map[string]interface{} `json:"payload"`
}

func (p *ForwardRequest) SendData() error {
	// TODO this needs to be pushed to the user?

	apiUrl := os.Getenv("AMAZON_LAMBDA_API_URL")

	// you cant just add the resource on the end of .com

	//TODO this resource needs to come from the database cardinal action
	resource := "/dev/handler/hello/"
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(p.Payload)

	u, _ := url.ParseRequestURI(apiUrl)
	u.Path = resource

	urlStr := fmt.Sprintf("%v", u) // "https://api.com/user/"

	client := &http.Client{}
	r, _ := http.NewRequest("POST", urlStr, b) // <-- URL-encoded payload
	r.Header.Add("Content-Type", "application/json")

	fmt.Println(b)

	resp, _ := client.Do(r)

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	jsonResponse := APIResponse{}

	jsonErr := json.Unmarshal(body, &jsonResponse)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	fmt.Println(jsonResponse)

	//respondWithJSON(w, http.StatusOK, jsonResponse)
	return nil
}