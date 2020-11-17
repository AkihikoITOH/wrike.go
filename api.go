package wrike

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/PuerkitoBio/httpcmw"
)

const (
	// APIRootPath is the API root path.
	APIRootPath = "api"
	// APIVersion is the API version number.
	APIVersion = "v4"
	// DefaultTimeout is the duration for which it waits for the API to respond.
	DefaultTimeout = 10 * time.Second
)

// API object contains the API configuration and has methods to communicate with the API.
type API struct {
	Config     *Config
	HTTPClient httpcmw.Doer
}

// NewAPI creates and returns a new API object based on the given Config.
func NewAPI(config *Config) *API {
	return &API{config, newHTTPClient()}
}

func (api *API) perform(req *http.Request) ([]byte, error) {
	res, err := api.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(res.Body)
}

func (api *API) get(basePath string, params *url.Values) ([]byte, error) {
	path := basePath
	if params != nil {
		path += fmt.Sprintf("?%s", params.Encode())
	}
	req, err := api.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	return api.perform(req)
}

func (api *API) put(path string, params *url.Values) ([]byte, error) {
	var data io.Reader
	if params != nil {
		data = strings.NewReader(params.Encode())
	} else {
		data = nil
	}
	req, err := api.newRequest(http.MethodPut, path, data)
	if err != nil {
		return nil, err
	}

	return api.perform(req)
}

func (api *API) post(path string, params *url.Values) ([]byte, error) {
	var data io.Reader
	if params != nil {
		data = strings.NewReader(params.Encode())
	} else {
		data = nil
	}
	req, err := api.newRequest(http.MethodPost, path, data)
	if err != nil {
		return nil, err
	}

	return api.perform(req)
}

func (api *API) delete(path string) ([]byte, error) {
	req, err := api.newRequest(http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	return api.perform(req)
}

func (api *API) newRequest(method, path string, body io.Reader) (*http.Request, error) {
	url := api.baseURL() + path
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("bearer %s", api.Config.APIAccessToken()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	return req, err
}

func (api *API) baseURL() string {
	return fmt.Sprintf("https://%s/%s/%s", api.Config.APIHost(), APIRootPath, APIVersion)
}

func newHTTPClient() *http.Client {
	client := &http.Client{Timeout: DefaultTimeout}
	return client
}
