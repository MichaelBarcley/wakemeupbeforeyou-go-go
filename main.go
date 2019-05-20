package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	url := "https://api.abiosgaming.com/v2/oauth/access_token"
	payload := strings.NewReader("grant_type=client_credentials&client_id=test-task&client_secret=9179d8d1b253209e193e7dee77e432ea79e541a5909a026a76")
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
	fmt.Fprintf(w, string(body))
}
