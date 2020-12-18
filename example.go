package gowally

import "github.com/Upper-Beacon/gowally/gohttp"

func exampleUsage() {
	client := gohttp.New()

	client.Get()
}
