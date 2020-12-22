package gowally

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Upper-Beacon/gowally/gohttp"
)

var (
	githubHTTPClient = getGithubClient()
)

func getGithubClient() gohttp.HTTPClient {
	client := gohttp.New()

	commomHeaders := make(http.Header)

	commomHeaders.Set("Authorization", "Bearer ABC-123")

	client.SetHeaders(commomHeaders)

	return client
}

func main() {
	getUrls()
}

func getUrls() {
	response, err := githubHTTPClient.Get("https://api.github.com", nil)

	if err != nil {
		panic(err)
	}

	fmt.Println(response.StatusCode)

	bytes, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(bytes))
}
