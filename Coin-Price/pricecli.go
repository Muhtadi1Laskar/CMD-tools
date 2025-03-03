package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func readJSON(url string, target any) error {
	response, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("unable to fetch data %v", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("error reading the response body: %v", err)
	}

	if err := json.Unmarshal(body, target); err != nil {
		return fmt.Errorf("error reading the response body: %v", err)
	}

	return nil
}

func main() {
	fmt.Println("Hello World")
}