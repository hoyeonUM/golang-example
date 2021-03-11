package exmaple_gorutine

type HttpClient interface {
	Request(url string, result chan HttpResponse)
}
type HttpResponse struct {
	StatusCode int
	Url        string
	Body       string

}

type Crawler struct {
	h HttpClient
}

func NewCrawler(h HttpClient) Crawler {
	return Crawler{h: h}
}

func (c *Crawler) DoAction(urls []string) []HttpResponse {
	ch := make(chan HttpResponse)
	for _, url := range urls {
		go c.h.Request(url, ch)
	}
	httpResponses := make([]HttpResponse, 0)
	for range urls {
		result := <-ch
		if result.StatusCode == 200 {
			httpResponses = append(httpResponses, result)
		}
	}
	return httpResponses
}
//mockgen . HttpClient > mock/http_client.go