package example_wrong_type_api_spec_test

import (
	"github.com/golang/mock/gomock"
	"github.com/hoyeonUM/golang-example/example_wrong_type_api_spec"
	mock_example_wrong_type_api_spec "github.com/hoyeonUM/golang-example/example_wrong_type_api_spec/mock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var response = `
{
	"name" : "test product",
	"price" : "",
	"created_at" : ""
}
`

func TestAPI(t *testing.T) {
	t.Run(`given wrong type response when use non safe struct json unmarshal then return error`, func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		m := mock_example_wrong_type_api_spec.NewMockHttpClient(ctrl)
		m.EXPECT().Request().Return(response).Times(1)
		api := example_wrong_type_api_spec.NewProductProxyAPI(m)
		_, err := api.Search()
		assert.Error(t, err)
	})

	t.Run(`given wrong type response when use safe struct json unmarshal then return struct and nil`, func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		m := mock_example_wrong_type_api_spec.NewMockHttpClient(ctrl)
		m.EXPECT().Request().Return(response).Times(1)
		api := example_wrong_type_api_spec.NewProductProxyAPI(m)
		result, err := api.SafeSearch()
		assert.Nil(t, err)
		assert.Equal(t, "test product", result.Name)
		assert.Equal(t, 0, result.Price.Value)
		assert.Equal(t, time.Time{}.Format(time.RFC3339), result.CreatedAt.Format(time.RFC3339))
	})
}
