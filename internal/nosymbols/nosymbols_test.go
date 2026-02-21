package nosymbols

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	
)

func TestNoSymbols(t *testing.T) {
	t.Run("invalid log message", func(t *testing.T) {
		message := "hello 🤠 world"

		err := validate(message)
		require.Error(t, err)
		assert.Equal(t, "special symbols or emojis are not allowed in logs", err.Error())
	})


	t.Run("valid log message", func(t *testing.T) {
		message := "hello world"

		err := validate(message)
		require.NoError(t, err)
	})

}
