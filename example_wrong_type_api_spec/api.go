package example_wrong_type_api_spec

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"
)

type ProductResponseSpec struct {
	Name      string    `json:"name"`
	Price     int       `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

type ProductResponseSpecProtect struct {
	Name      string   `json:"name"`
	Price     SafeInt  `json:"price"`
	CreatedAt SafeTime `json:"created_at"`
}

type SafeTime struct {
	time.Time
}
type SafeInt struct {
	Value int
}

func (m *SafeInt) UnmarshalJSON(data []byte) error {
	strData := string(data)
	if strData == "null" || strData == `""` {
		return nil
	}
	n, err := strconv.Atoi(strings.ReplaceAll(strData, "\"", ""))
	*m = SafeInt{n}
	return err
}

func (m *SafeTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" || string(data) == `""` {
		return nil
	}
	tt, err := time.Parse(`"`+time.RFC3339+`"`, string(data))
	*m = SafeTime{tt}
	return err
}

type HttpClient interface {
	Request() string
}
type HttpResponse struct {
	StatusCode int
	Url        string
	Body       string
}

type ProductProxyAPI struct {
	h HttpClient
}

func NewProductProxyAPI(h HttpClient) ProductProxyAPI {
	return ProductProxyAPI{h: h}
}

func (p ProductProxyAPI) Search() (ProductResponseSpec, error) {
	response := p.h.Request()
	spec := ProductResponseSpec{}
	if err := json.Unmarshal([]byte(response), &spec); err != nil {
		return ProductResponseSpec{}, err
	}
	return spec, nil
}

func (p ProductProxyAPI) SafeSearch() (ProductResponseSpecProtect, error) {
	response := p.h.Request()
	spec := ProductResponseSpecProtect{}
	if err := json.Unmarshal([]byte(response), &spec); err != nil {
		return ProductResponseSpecProtect{}, err
	}
	return spec, nil
}
