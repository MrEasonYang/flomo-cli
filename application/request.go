package application

import (
	"net/http"
	"encoding/json"
	"bytes"
)

// SendMemo is the func to request flomo server and save the content.
func SendMemo(content string) {
	if content == "" {
		panic("The memo content is empty, type something and try it again.")
	}
	config := GetConfig()	
	if config.Api == "" {
		panic("Flomo API is not set.")
	}

	body := ReqBody{}
	body.Content = content
	jsonValue, err := json.Marshal(body)
	if err != nil {
		panic("Failed to parse flomo memo request.")
	}

	resp, err := http.Post(
		config.Api, 
		"application/json; charset=utf-8", 
		bytes.NewBuffer(jsonValue))
	if err != nil {
		panic(err)
	}
	if resp.StatusCode != 200 {
		panic("Flomo server responsed with unsuccessful result")
	}
	defer resp.Body.Close()
}

// ReqBody is the struct of the HTTP request for flomo.
type ReqBody struct {
	Content string `json:"content"`
}