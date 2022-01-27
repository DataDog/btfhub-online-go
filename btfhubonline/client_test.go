package btfhubonline

import (
	"testing"
)

func TestNewClient(t *testing.T) {
	tests := []struct {
		name          string
		serverAddress string
		wantErr       bool
	}{
		{
			name:          "url without scheme",
			serverAddress: "btfhub.seekret.com",
			wantErr:       true,
		},
		{
			name:          "ip without scheme",
			serverAddress: "192.168.0.1",
			wantErr:       true,
		},

		{
			name:          "http url with scheme",
			serverAddress: "http://btfhub.seekret.com",
			wantErr:       false,
		},
		{
			name:          "https url with scheme",
			serverAddress: "https://btfhub.seekret.com",
			wantErr:       false,
		},
		{
			name:          "http ip with scheme",
			serverAddress: "http://192.168.0.1",
			wantErr:       false,
		},
		{
			name:          "https ip with scheme",
			serverAddress: "https://192.168.0.1",
			wantErr:       false,
		},

		{
			name:          "invalid url",
			serverAddress: "https://btfhub .seekret.com",
			wantErr:       true,
		},

		{
			name:          "http url with port",
			serverAddress: "http://btfhub.seekret.com:8080",
			wantErr:       false,
		},
		{
			name:          "https url with port",
			serverAddress: "https://btfhub.seekret.com:8080",
			wantErr:       false,
		},
		{
			name:          "http ip with port",
			serverAddress: "http://192.168.0.1:8080",
			wantErr:       false,
		},
		{
			name:          "https ip with port",
			serverAddress: "https://192.168.0.1:8080",
			wantErr:       false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := New(tt.serverAddress)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
