package common

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindLineNo(t *testing.T) {
	t.Run("Returns error if value is not present - empty lines", func(t *testing.T) {
		var lines []*string
		yamlPath := "a.b"

		res, err := FindLineNo(lines,  yamlPath)

		assert.Nil(t, res)
		assert.Error(t, err)
	})

	t.Run("Returns error if value is not present", func(t *testing.T) {
		line1 := "a:"
		line2 := "  c"
		lines := []*string{
			&line1,
			&line2,
		}
		yamlPath := "a.b"

		res, err := FindLineNo(lines,  yamlPath)

		assert.Nil(t, res)
		assert.Error(t, err)
	})

	t.Run("Returns line number if value is present", func(t *testing.T) {
		line1 := "a:"
		line2 := "  b: test"
		lines := []*string{
			&line1,
			&line2,
		}
		yamlPath := "a.b"

		res, err := FindLineNo(lines,  yamlPath)

		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, *res, 1)
	})

	t.Run("A lil more complex case", func(t *testing.T) {
		line1 := "b:"
		line2 := "  c: test"
		line3 := "a:"
		line4 := "  b:"
		line5 := "    c: test"
		line6 := "  c: test"
		lines := []*string{
			&line1,
			&line2,
			&line3,
			&line4,
			&line5,
			&line6,
		}
		yamlPath := "a.b.c"

		res, err := FindLineNo(lines,  yamlPath)

		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, *res, 4)
	})
}
