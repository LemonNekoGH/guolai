package guolai

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
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

// makeRequest create and make a request to wolai api.
//
// Params:
// - path: should start with `/`
func makeRequest(
	path string,
	method string,
	token string,
	httpClient HTTPClient,
	body any,
	dataType any,
) (*WolaiResponse, error) {
	client := httpClient
	if client == nil {
		client = http.DefaultClient
	}

	var (
		bodyJson []byte
		err      error
	)
	if body != nil {
		bodyJson, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, "https://openapi.wolai.com"+path, bytes.NewReader(bodyJson))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	if token != "" {
		req.Header.Set("Authorization", token)
	}

	resp, err := client.Do(req)
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

	if resp.StatusCode >= http.StatusBadRequest {
		return respStruct, errors.New(respStruct.Message)
	}

	return respStruct, nil
}

// GetBlocks https://www.wolai.com/wolai/htv4BHWQwiqfqx6PHQUrRn
func (api *WolaiAPI) GetBlocks(blockId string) (*BlockApiResponse, error) {
	resp, err := makeRequest(
		"/v1/blocks/"+blockId,
		http.MethodGet,
		api.Token,
		api.HTTPClient,
		nil,
		&BlockApiResponse{},
	)
	if err != nil {
		return nil, err
	}

	respData := resp.Data.(*BlockApiResponse)

	return respData, nil
}

func (api *WolaiAPI) GetBlockChildren(blockId string) ([]BlockApiResponse, error) {
	resp, err := makeRequest(
		"/v1/blocks/"+blockId+"/children",
		http.MethodGet,
		api.Token,
		api.HTTPClient,
		nil,
		&[]BlockApiResponse{},
	)
	if err != nil {
		return nil, err
	}

	respData := resp.Data.(*[]BlockApiResponse)

	return *respData, nil
}

// CreateBlocks https://www.wolai.com/wolai/oyKuZbAmufkA3r7ocrBxW2
func (api *WolaiAPI) CreateBlocks(parentId string, blocks []Block) ([]string, error) {
	resp, err := makeRequest(
		"/v1/blocks",
		http.MethodPost,
		api.Token,
		api.HTTPClient,
		map[string]any{
			"parent_id": parentId,
			"blocks":    blocks,
		},
		&[]string{},
	)
	if err != nil {
		return []string{}, err
	}

	respData := resp.Data.(*[]string)

	return *respData, nil
}

// TODO: helper functions for creating blocks
