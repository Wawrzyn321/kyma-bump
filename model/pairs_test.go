package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPairs(t *testing.T) {
	mappings := Mappings{
		Mapping{
			Name:        "name",
			Aliases:     []string{"name", "n"},
		},
		Mapping{
			Name:        "other",
			Aliases:     []string{"other"},
		},
	}

	t.Run("Resolves single alias", func(t *testing.T) {
		pairs := PairCollection{
			"n": "name-tag",
		}

		resolvedPairs, errs := pairs.Dealiasize(mappings)

		assert.Len(t, errs, 0)
		assert.Len(t, resolvedPairs, 1)
		assert.Equal(t, resolvedPairs["name"], "name-tag")
	})

	t.Run("Resolves multiple aliases", func(t *testing.T) {
		pairs := PairCollection{
			"n": "name-tag",
			"other": "other-tag",
		}

		resolvedPairs, errs := pairs.Dealiasize(mappings)

		assert.Len(t, errs, 0)
		assert.Len(t, resolvedPairs, 2)
		assert.Equal(t, resolvedPairs["name"], "name-tag")
		assert.Equal(t, resolvedPairs["other"], "other-tag")
	})

	t.Run("Aggregates errors", func(t *testing.T) {
		pairs := PairCollection{
			"a": "name-tag",
			"b": "other-tag",
		}

		resolvedPairs, errs := pairs.Dealiasize(mappings)

		assert.Len(t, errs, 2)
		assert.Len(t, resolvedPairs, 0)
	})
}
