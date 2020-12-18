package gowally

import (
	"fmt"
	"net/http"

	"github.com/Upper-Beacon/gowally/gohttp"
)

func basicExample() {
	client := gohttp.New()

	headers := make(http.Header)

	headers.Set("Authorization", "Bearer iuashdiushiudhasiudhaisd")

	response, err := client.Get("https://api.github.com", headers)

	if err != nil {
		panic(err)
	}

	fmt.Sprintln(response.StatusCode)
}
