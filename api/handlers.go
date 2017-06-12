package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type gitHubCodeQuery struct {
	Code string `json:"code"`
}

type gitHubTokenQuery struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Code         string `json:"code"`
}

// Endpoint test
func Endpoint(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is a cool message!"))
}

// GithubCode get github code
func GithubCode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	q := &gitHubCodeQuery{}
	if err := json.NewDecoder(r.Body).Decode(q); err != nil {
		log.Printf("Error decoding body: %s", err)
		return
	}
	fmt.Println(q.Code)

	tokenQuery := gitHubTokenQuery{
		ClientID:     "95373b61a45242af69d2",
		ClientSecret: "f28c3097d688c6270889761c233169f29d78199c",
		Code:         q.Code,
	}

	getGitHubToken(&tokenQuery)

	res := map[string]string{
		"message": "Well Done !",
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		panic(err)
	}
}

func getGitHubToken(q *gitHubTokenQuery) string {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(q)

	resp, _ := http.Post(
		"https://github.com/login/oauth/access_token",
		"application/json; charset=utf-8",
		b,
	)

	// read the response body to a variable
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)
	//print raw response body for debugging purposes
	fmt.Println("\n\n", bodyString)

	return "aa"
}
