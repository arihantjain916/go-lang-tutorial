package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://dogapi.dog/api/v1/facts"

func main() {

	fmt.Println("Welcome to Web Folder ")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)

	fmt.Println("Body is :", string(body))
}
