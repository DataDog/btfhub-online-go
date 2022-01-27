package btfhubonline

import (
	"fmt"
	"net/url"
	"strings"
)

const (
	httpPrefix  = "http://"
	httpsPrefix = "https://"
)

// craftURL builds a valid URL out of serverAddress and the secure flag.
func craftURL(serverAddress string, secure bool) (string, error) {
	if strings.HasPrefix(serverAddress, httpPrefix) || strings.HasPrefix(serverAddress, httpsPrefix) {
		return "", fmt.Errorf("server address %q should not hold a URL scheme, please use ClientOptions instead", serverAddress)
	}
	scheme := httpPrefix
	if secure {
		scheme = httpsPrefix
	}

	_, err := url.ParseRequestURI(scheme + serverAddress)
	if err != nil {
		return "", err
	}
	return scheme + serverAddress, nil
}
