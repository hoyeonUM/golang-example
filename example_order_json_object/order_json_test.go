package example_order_json_object_test

import (
	"encoding/json"
	"github.com/hoyeonUM/golang-example/example_order_json_object"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOrderJson(t *testing.T) {
	t.Run(`given string map first key is b and second key is a when parse json then first key is a and second key is b`, func(t *testing.T) {
		notOrderJson := make(map[string]string)
		notOrderJson["b"] = "bbb"
		notOrderJson["a"] = "aaa"
		b, _ := json.Marshal(notOrderJson)
		assert.Equal(t, `{"a":"aaa","b":"bbb"}`, string(b))
	})

	t.Run(`given string map first key is b and second key is a when parse order json then first key is b and second key is a`, func(t *testing.T) {
		notOrderJson := make(map[string]string)
		notOrderJson["b"] = "bbb"
		notOrderJson["a"] = "aaa"

		OrderJson := example_order_json_object.OrderJsonMap{}
		OrderJson.Order = []string{"b", "a"}
		OrderJson.Map = notOrderJson
		b, _ := json.Marshal(OrderJson)
		assert.Equal(t, `{"b":"bbb","a":"aaa"}`, string(b))
	})
}
