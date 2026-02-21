package sensitivity

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSensivity(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		patterns []string
		wantErr  bool
	}{
		{"password", "my_password_123", []string{"password", "secret", "token", "auth", "apiKey", "api_key"}, true},
		{"secret", "super_secret_value", []string{"password", "secret", "token", "auth", "apiKey", "api_key"}, true},
		{"token", "access_token_abc", []string{"password", "secret", "token", "auth", "apiKey", "api_key"}, true},
		{"api_key", "api_key=asdasda", []string{"password", "secret", "token", "auth", "apiKey", "api_key"}, true},

		{"valid message", "hello world", []string{"password", "secret", "token", "auth", "apiKey", "api_key"}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validate(tt.input, tt.patterns)

			if tt.wantErr == true {
				require.Error(t, err, "potential sensitive data leak", tt.input)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
