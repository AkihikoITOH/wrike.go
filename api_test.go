package wrike_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"

	"gotest.tools/assert"

	wrike "github.com/AkihikoITOH/wrike.go"
)

const (
	apiAccessToken = "wrike-access-token"
)

var (
	mockAPIConfig = wrike.NewConfig(apiAccessToken, "")
)

type MockHTTPClient struct {
	Request  *http.Request
	Response *http.Response
}

func (mockClient *MockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	mockClient.Request = req
	return mockClient.Response, nil
}

func NewMockClient(body []byte) *MockHTTPClient {
	responseHeader := http.Header{}
	responseHeader.Add("Date", "Fri, 08 Mar 2019 19:36:47 GMT")
	responseHeader.Add("Content-Type", "application/json;charset=UTF-8")
	responseBody := ioutil.NopCloser(bytes.NewReader(body))
	mockResponse := &http.Response{
		Status:        "200 OK",
		StatusCode:    200,
		Proto:         "HTTP/2.0",
		ProtoMajor:    2,
		ProtoMinor:    0,
		Header:        responseHeader,
		Body:          responseBody,
		ContentLength: -1,
		Close:         false,
		Uncompressed:  false,
	}
	return &MockHTTPClient{Response: mockResponse}
}

func SharedRequestTests(t *testing.T, req *http.Request) {
	assert.Equal(t, req.Host, wrike.DefaultAPIHost)
	header := req.Header
	assert.Equal(t, header["Authorization"][0], "bearer wrike-access-token")
}

func TestNewAPI(t *testing.T) {
	config := &wrike.Config{}
	api := wrike.NewAPI(config)

	assert.Equal(t, reflect.TypeOf(api).String(), "*wrike.API")
}
