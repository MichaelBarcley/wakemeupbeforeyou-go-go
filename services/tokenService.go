package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

// AccessToken is used to access the Abios Gaming API. On server start it is empty.
var AccessToken = ""

// TokenCreationDate is the date when the AccessToken was created. On server start it is time.Now() and it is modified in token handling functions.
var TokenCreationDate = time.Now()


// GetToken ...
func GetToken() {
	url := "https://api.abiosgaming.com/v2/oauth/access_token"
	payload := strings.NewReader("grant_type=client_credentials&client_id=test-task&client_secret=9179d8d1b253209e193e7dee77e432ea79e541a5909a026a76")
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))

	m := make(map[string]interface{})
	json.Unmarshal(body, &m)
	fmt.Println("Your access token for the next hour is: ", m["access_token"])
	AccessToken = m["access_token"].(string)
	TokenCreationDate = time.Now()
}

// CheckIfTokenIsValid ...
func CheckIfTokenIsValid() {
	var timeSinceTokenCreation = time.Now().Sub(TokenCreationDate) / 10e8
	fmt.Println("The age of the token is: ", timeSinceTokenCreation)
	if timeSinceTokenCreation > 3600 || AccessToken == "" {
		GetToken()
	}
}
