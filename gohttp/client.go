package gohttp

type httpClient struct{}

// HTTPClient ...
type HTTPClient interface {
	Get()
	Post()
	Put()
	Patch()
	Delete()
}

func (c *httpClient) Get() {}

func (c *httpClient) Post() {}

func (c *httpClient) Put() {}

func (c *httpClient) Patch() {}

func (c *httpClient) Delete() {}
