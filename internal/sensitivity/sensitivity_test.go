package sensitivity

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSensivity(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{"password", "my_password_123", true},
		{"secret", "super_secret_value", true},
		{"token", "access_token_abc", true},
		{"api_key", "api_key=asdasda", true},

		{"valid message", "hello world", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validate(tt.input)

			if tt.wantErr == true {
				require.Error(t, err, "potential sensitive data leak", tt.input)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
