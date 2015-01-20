package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "http://127.0.0.1:9000/hello"
	_, err := getJSON(url)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func getJSON(url string) ([]byte, error) {
	request, err := http.NewRequest("Get", url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add("Accept", "application/json")
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	jsonContent, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	} else {
		fmt.Println(string(jsonContent))
		return jsonContent, nil
	}
}
