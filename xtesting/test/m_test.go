package do_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTest(t *testing.T) {
	t.Run("check nil", func(t *testing.T) {
		var expectedErr error = nil
		assert.Equal(t, expectedErr, nil)
	})
}
