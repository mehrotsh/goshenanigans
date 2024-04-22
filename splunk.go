package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Payload struct {
	Event string `json:"event"`
}

func main() {
	url := "https://<splunk-instance>:8088/services/collector/raw"
	token := "<token>"

	newStr := `This is a "string" with "double quotes"`

	data := Payload{
		Event: newStr,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Splunk "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Print the HTTP response status.
	fmt.Println("Response Status:", resp.Status)
}
