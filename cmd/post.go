package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	contact := `{"name":"Petur", "address":"Pleven", "age":11, "email":"pepi@test.com"}`

	resp, err := http.Post("http://localhost:8080/contacts", "application/json", strings.NewReader(contact))
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(body))
}
