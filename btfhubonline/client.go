package btfhubonline

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/go-resty/resty/v2"
	"github.com/seek-ret/btfhub-online-go/internal/compression"
)

// Client exposes all API available in the BTFHub online server.
type Client struct {
	serverAddress string
}

// New initiates a new instance of BTFHub online client.
func New(serverAddress string) (Client, error) {
	if _, err := url.ParseRequestURI(serverAddress); err != nil {
		return Client{}, fmt.Errorf("url %q is invalid: %v", serverAddress, err)
	}
	return Client{
		serverAddress: serverAddress,
	}, nil
}

// getBaseRequest returns a new request instance with all common headers.
func getBaseRequest() *resty.Request {
	request := resty.New().R()
	request.SetHeader("User-Agent", fmt.Sprintf("btfhub-online-go/%s", Version))
	return request
}

// List returns a list of BTFRecordIdentifier objects that exist in the server.
func (client Client) List() ([]BTFRecordIdentifier, error) {
	request := getBaseRequest()
	response, err := request.Get(fmt.Sprintf("%s/api/v1/list", client.serverAddress))
	if err != nil {
		return nil, err
	}
	// IsError is not the opposite of IsSuccess.
	if !response.IsSuccess() {
		return nil, fmt.Errorf("failed to retrieve BTFs list. Detailed message: Response code: %d, message: %s", response.StatusCode(), response.String())
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
		return nil, fmt.Errorf("failed to retrieve BTF. Detailed message: Response code: %d. message: %s", response.StatusCode(), response.String())
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
		return nil, fmt.Errorf("failed to retrieve a custom BTF. Detailed message: Response code %d, message: %s", response.StatusCode(), response.String())
	}
	return compression.ExtractFileTarXZ(bytes.NewReader(response.Body()))
}
