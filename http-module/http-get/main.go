package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	url := "https://jsonplaceholder.typicode.com/todos"

	res, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer res.Body.Close()

	fmt.Println("Status Code:", res.StatusCode)
	fmt.Println("Status :", res.Status)

	if res.StatusCode != http.StatusOK {
		fmt.Println(res.StatusCode)
		return  
	}

	bodyBytes, err := io.ReadAll(res.Body)

	if err != nil {
		fmt.Println(err)
		return
	}

	bodyText := string(bodyBytes)

	max := 250

	if len(bodyText) < max {
		max = len(bodyText)
	}

	fmt.Println(bodyText[:max])
}