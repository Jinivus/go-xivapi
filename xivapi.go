package xivapi

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	defaultBaseURL = "https://xivapi.com/"
)

var errNonNilContext = errors.New("context must be non-nil")

type Client struct {
	client *http.Client

	// Base URL with trailing slash, defaults to "https://xivapi.com/"
	BaseUrl *url.URL

	common service

	Search *SearchService
}

type service struct {
	client *Client
}

func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	baseUrl, _ := url.Parse(defaultBaseURL)
	c := &Client{client: httpClient, BaseUrl: baseUrl}
	c.common.client = c
	c.Search = (*SearchService)(&c.common)

	return c
}

type Pagination struct {
	Page         *int `json:"Page,omitempty"`
	PageNext     *int `json:"PageNext,omitempty"`
	PagePrev     *int `json:"PagePrev,omitempty"`
	Results      *int `json:"Results,omitempty"`
	ResultsTotal *int `json:"ResultsTotal,omitempty"`
}

type PaginatedResult struct {
	Pagination *Pagination `json:"Pagination,omitempty"`
}

func (c *Client) NewRequest(method, urlString string, body interface{}) (*http.Request, error) {
	if !strings.HasSuffix(c.BaseUrl.Path, "/") {
		return nil, fmt.Errorf("BaseURL must have a trailing slash, but %q does not", c.BaseUrl)
	}

	u, err := c.BaseUrl.Parse(urlString)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf = &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	return req, nil
}

func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*Response, error) {

	bareResp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	resp := newResponse(bareResp)
	defer resp.Body.Close()

	switch v := v.(type) {
	case nil:
	case io.Writer:
		_, err = io.Copy(v, resp.Body)
	default:
		decErr := json.NewDecoder(resp.Body).Decode(v)
		if decErr == io.EOF {
			decErr = nil // ignore EOF errors caused by empty response body
		}
		if decErr != nil {
			err = decErr
		}
	}
	return resp, err
}

func newResponse(r *http.Response) *Response {
	response := &Response{Response: r}
	response.populatePageValues()
	return response
}

// populatePageValues parses the HTTP Link response headers and populates the
// various pagination link values in the Response.
func (r *Response) populatePageValues() {
	//TODO: Populate page values from
	// "Pagination": {
	//	"Page": 23,
	//	"PageNext": null,
	//	"PagePrev": 22,
	//	"PageTotal": 23,
	//	"Results": 44,
	//	"ResultsPerPage": 100,
	//	"ResultsTotal": 2244
	// },
}

type Response struct {
	*http.Response

	NextPage int
	PrevPage int
	CurPage  int
}
