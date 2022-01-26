package btfhubonline

import (
	"reflect"
	"testing"
)

func TestCraftURL(t *testing.T) {
	tests := []struct {
		name          string
		serverAddress string
		secure        bool
		want          string
		wantErr       bool
	}{
		{
			name:          "http url",
			serverAddress: "btfhub.seekret.com",
			secure:        false,
			want:          "http://btfhub.seekret.com",
			wantErr:       false,
		},
		{
			name:          "https url",
			serverAddress: "btfhub.seekret.com",
			secure:        true,
			want:          "https://btfhub.seekret.com",
			wantErr:       false,
		},
		{
			name:          "http ip",
			serverAddress: "192.168.0.1",
			secure:        false,
			want:          "http://192.168.0.1",
			wantErr:       false,
		},
		{
			name:          "https ip",
			serverAddress: "192.168.0.1",
			secure:        true,
			want:          "https://192.168.0.1",
			wantErr:       false,
		},

		{
			name:          "http url with scheme",
			serverAddress: "http://btfhub.seekret.com",
			secure:        false,
			want:          "",
			wantErr:       true,
		},
		{
			name:          "https url with scheme",
			serverAddress: "https://btfhub.seekret.com",
			secure:        true,
			want:          "",
			wantErr:       true,
		},
		{
			name:          "http ip with scheme",
			serverAddress: "http://192.168.0.1",
			secure:        false,
			want:          "",
			wantErr:       true,
		},
		{
			name:          "https ip with scheme",
			serverAddress: "https://192.168.0.1",
			secure:        true,
			want:          "",
			wantErr:       true,
		},

		{
			name:          "invalid url",
			serverAddress: "btfhub .seekret.com",
			want:          "",
			wantErr:       true,
		},

		{
			name:          "http url with port",
			serverAddress: "btfhub.seekret.com:8080",
			secure:        false,
			want:          "http://btfhub.seekret.com:8080",
			wantErr:       false,
		},
		{
			name:          "https url with port",
			serverAddress: "btfhub.seekret.com:8080",
			secure:        true,
			want:          "https://btfhub.seekret.com:8080",
			wantErr:       false,
		},
		{
			name:          "http ip with port",
			serverAddress: "192.168.0.1:8080",
			secure:        false,
			want:          "http://192.168.0.1:8080",
			wantErr:       false,
		},
		{
			name:          "https ip with port",
			serverAddress: "192.168.0.1:8080",
			secure:        true,
			want:          "https://192.168.0.1:8080",
			wantErr:       false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := craftURL(tt.serverAddress, tt.secure)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() got = %v, want %v", got, tt.want)
			}
		})
	}
}
