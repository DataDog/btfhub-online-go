package btfhubonline

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
)

// Client exposes all apis available in the BTFHub online server.
type Client struct {
	serverAddress string
}

// ClientOptions holds the different options for connecting to the client.
type ClientOptions struct {
	Secure bool
}

// New initiates a new instance of BTFHub online client.
func New(serverAddress string, opts ClientOptions) (Client, error) {
	finalURL, err := craftURL(serverAddress, opts.Secure)
	if err != nil {
		return Client{}, fmt.Errorf("failed crafting url due to: %v", err)
	}
	return Client{
		serverAddress: finalURL,
	}, nil
}

// getBaseRequest returns a new request instance with all common headers.
func getBaseRequest() *resty.Request {
	request := resty.New().R()
	request.SetHeader("User-Agent", fmt.Sprintf("btfhub-online-go/%s", Version))
	return request
}

// List returns a list of BTFRecordIdentifier that exist in the server.
func (client Client) List() ([]BTFRecordIdentifier, error) {
	request := getBaseRequest()
	response, err := request.Get(fmt.Sprintf("%s/api/v1/list", client.serverAddress))
	if err != nil {
		return nil, err
	}
	// IsError is not the opposite of IsSuccess.
	if !response.IsSuccess() {
		return nil, fmt.Errorf("expeted to get a success status code, but instead got %d. Error: %s", response.StatusCode(), response.String())
	}

	var btfList []BTFRecordIdentifier
	if err := json.Unmarshal(response.Body(), &btfList); err != nil {
		return nil, fmt.Errorf("failed unmarshaling BTF list from response. Error: %s", err)
	}

	return btfList, nil
}

// GetRawBTF returns a BTF from the server without customizing it for a specific BPF.
func (client Client) GetRawBTF(btfIdentifier BTFRecordIdentifier) ([]byte, error) {
	request := getBaseRequest()
	request.SetQueryParams(btfIdentifier.AsMap())
	response, err := request.Get(fmt.Sprintf("%s/api/v1/download", client.serverAddress))
	if err != nil {
		return nil, err
	}
	// IsError is not the opposite of IsSuccess.
	if !response.IsSuccess() {
		return nil, fmt.Errorf("expeted to get a success status code, but instead got %d. Error: %s", response.StatusCode(), string(response.Body()))
	}
	return response.Body(), nil
}

// GetCustomBTF returns a BTF from the server with customizing it for a specific BPF.
func (client Client) GetCustomBTF(btfIdentifier BTFRecordIdentifier, bpfBinary []byte) ([]byte, error) {
	request := getBaseRequest()
	request.SetQueryParams(btfIdentifier.AsMap())
	request.SetFileReader("bpf", "from_memory", bytes.NewReader(bpfBinary))
	response, err := request.Post(fmt.Sprintf("%s/api/v1/customize", client.serverAddress))
	if err != nil {
		return nil, err
	}
	// IsError is not the opposite of IsSuccess.
	if !response.IsSuccess() {
		return nil, fmt.Errorf("expeted to get a success status code, but instead got %d. Error: %s", response.StatusCode(), string(response.Body()))
	}
	return response.Body(), nil
}
