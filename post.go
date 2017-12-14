package main

import (
	"fmt"
	"net/http"
)

func Init() {
	client := http.Client{}
	res, err := client.Get("https://www.google.com")
	if err != nil {
		fmt.Println("Init Error:", err)
	}
	res.Body.Close()

}
