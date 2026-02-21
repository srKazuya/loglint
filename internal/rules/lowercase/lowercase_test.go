package lowercase

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	
)

func TestLowerCase(t *testing.T) {
	t.Run("invalid log message", func(t *testing.T) {
		message := "Hello world"

		err := validate(message)
		require.Error(t, err)
		assert.Equal(t, "log message must begin with a lowercase letter", err.Error())
	})


	t.Run("valid log message", func(t *testing.T) {
		message := "hello world"

		err := validate(message)
		require.NoError(t, err)
	})

}
