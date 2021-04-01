package example_gorutine_test

import (
	"github.com/golang/mock/gomock"
	"github.com/hoyeonUM/golang-example/example_gorutine"
	mock_exmaple "github.com/hoyeonUM/golang-example/example_gorutine/mock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCrawler(t *testing.T) {
	t.Run("Given three crawling target When one fails Then Only successful responses are returned", func(t *testing.T) {
		urls := []string{"https://www.naver.com", "https://google.com", "https://www.cafe24.com"}
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		m := mock_exmaple.NewMockHttpClient(ctrl)
		m.EXPECT().Request(gomock.Any(), gomock.Any()).
			Do(func(arg0, arg1 interface{}) {
				url := arg0.(string)
				response := example_gorutine.HttpResponse{
					Body:       "<html></html>",
					StatusCode: 200,
					Url:        url,
				}
				if url == "https://www.naver.com" {
					response.StatusCode = 404
				}
				arg1.(chan example_gorutine.HttpResponse) <- response
			}).
			Times(len(urls))
		c := example_gorutine.NewCrawler(m)
		responses := c.DoAction(urls)
		assert.Len(t, responses, 2)
		successUrls := make([]string, 0)
		for _, response := range responses {
			successUrls = append(successUrls, response.Url)
		}
		assert.Contains(t, successUrls, "https://google.com")
		assert.Contains(t, successUrls, "https://www.cafe24.com")
		assert.NotContains(t, successUrls, "https://www.naver.com")
	})
}

//mockgen . HttpClient > mock\http_client.go
