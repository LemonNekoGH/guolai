package wolaiapi

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
)

// HTTPClient is the type needed for the bot to perform HTTP requests.
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type WolaiAPI struct {
	Token      string
	HTTPClient HTTPClient
}

func New(token string) *WolaiAPI {
	return NewWithClient(token, &http.Client{})
}

func NewWithClient(token string, httpClient HTTPClient) *WolaiAPI {
	return &WolaiAPI{
		Token:      token,
		HTTPClient: httpClient,
	}
}

// MakeRequest create and make a request to wolai api.
//
// Params:
// - path: should start with `/`
func (api *WolaiAPI) MakeRequest(path string, method string, body any, dataType any) (*WolaiResponse, error) {
	reqUrl, err := url.Parse("https://openapi.wolai.com" + path)
	if err != nil {
		return nil, err
	}

	var bodyJson []byte
	var bodyReader io.ReadCloser
	if body != nil {
		bodyJson, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}

		bodyReader = io.NopCloser(bytes.NewReader(bodyJson))
	}

	req := &http.Request{
		Method: method,
		URL:    reqUrl,
		Body:   bodyReader,
		Header: map[string][]string{
			"Accept":        {"application/json"},
			"Authorization": {api.Token},
			"Content-Type":  {"application/json"},
		},
	}

	resp, err := api.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer func(body io.ReadCloser) {
		err := body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	respStruct := &WolaiResponse{
		Data: dataType,
	}
	err = json.Unmarshal(respBody, respStruct)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return respStruct, errors.New(respStruct.Message)
	}

	return respStruct, nil
}

// GetBlocks https://www.wolai.com/wolai/htv4BHWQwiqfqx6PHQUrRn
func (api *WolaiAPI) GetBlocks(blockId string) (*BlockApiResponse, error) {
	resp, err := api.MakeRequest("/v1/blocks/"+blockId, http.MethodGet, nil, &BlockApiResponse{})
	if err != nil {
		return nil, err
	}

	respData := resp.Data.(*BlockApiResponse)

	return respData, nil
}

func (api *WolaiAPI) GetBlockChildren(blockId string) ([]BlockApiResponse, error) {
	resp, err := api.MakeRequest("/v1/blocks/"+blockId+"/children", http.MethodGet, nil, &[]BlockApiResponse{})
	if err != nil {
		return nil, err
	}

	respData := resp.Data.(*[]BlockApiResponse)

	return *respData, nil
}
