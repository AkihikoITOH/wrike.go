package wrike

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
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

func (api *API) get(path string) ([]byte, error) {
	req, err := api.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	res, err := api.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(res.Body)
}

func (api *API) post(path string, body []byte) ([]byte, error) {
	return nil, fmt.Errorf("method %s is not implemented yet", http.MethodPost)
}

func (api *API) delete(path string) ([]byte, error) {
	return nil, fmt.Errorf("method %s is not implemented yet", http.MethodDelete)
}

func (api *API) newRequest(method, path string, body io.Reader) (*http.Request, error) {
	url := api.baseURL() + path
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("bearer %s", api.Config.APIAccessToken()))

	return req, err
}

func (api *API) baseURL() string {
	return fmt.Sprintf("https://%s/%s/%s", api.Config.APIHost(), APIRootPath, APIVersion)
}

func newHTTPClient() *http.Client {
	client := &http.Client{Timeout: DefaultTimeout}
	return client
}
