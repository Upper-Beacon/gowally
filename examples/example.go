package gowally

import (
	"fmt"

	"github.com/Upper-Beacon/gowally/gohttp"
)

func basicExample() {
	client := gohttp.New()

	response, err := client.Get("https://api.github.com", nil)

	if err != nil {
		panic(err)
	}

	fmt.Sprintln(response.StatusCode)
}
