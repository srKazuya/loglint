package englishonly

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	
)

func TestEnglishOnly(t *testing.T) {
	t.Run("invalid log message", func(t *testing.T) {
		message := "hello мир"

		err := validate(message)
		require.Error(t, err)
		assert.Equal(t, "non-english characters detected (use latin alphabet only)", err.Error())
	})


	t.Run("valid log message", func(t *testing.T) {
		message := "hello world"

		err := validate(message)
		require.NoError(t, err)
	})

}
