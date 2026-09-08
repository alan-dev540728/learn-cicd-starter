package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		Authorization     string
		ApiKey            string
		WantAuthorization string
		WantApiKey        string
		expectedErr       bool
		name              string
	}{
		{
			Authorization:     "Authorization",
			ApiKey:            "ApiKey 123456",
			WantAuthorization: "Authorization",
			WantApiKey:        "123456",
			expectedErr:       false,
		},
		{
			Authorization:     "",
			ApiKey:            "ApiKey 123456",
			WantAuthorization: "Authorization",
			WantApiKey:        "456789",
			expectedErr:       true,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			headers := http.Header{}
			if tc.Authorization != "" {
				headers.Set(tc.Authorization, tc.ApiKey)
			}
			got, err := GetAPIKey(headers)

			if (err != nil) != tc.expectedErr {
				t.Fatalf("expected error: %v, got: %v", tc.expectedErr, err)
			}
			if !tc.expectedErr && got != tc.WantApiKey {
				t.Errorf("got %q, want %q", got, tc.WantApiKey)
			}

		})
	}

}
