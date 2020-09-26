package model

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestPair struct {
	Name string
	ResolvedName *string
}

func TestMappings_ResolveName(t *testing.T) {
	m := mockMappings()

	resolvedName := "name"
	testPairs := []TestPair{
		{
			Name:         "name",
			ResolvedName: &resolvedName,
		},
		{
			Name:         "n",
			ResolvedName: &resolvedName,
		},
		{
			Name:         "o",
			ResolvedName: nil,
		},
	}

	for _, pair := range testPairs {
		t.Run(fmt.Sprintf("Resolves name by alias - %s", pair.Name), func(t *testing.T) {
			resolvedName := m.ResolveName(pair.Name)

			assert.Equal(t, pair.ResolvedName, resolvedName)
		})
	}
}

func TestMappings_FindByName(t *testing.T) {
	m := mockMappings()

	t.Run("Finds by name, returning mapping if exists", func(t *testing.T) {
		mapping := m.FindByName("other")

		assert.NotNil(t, mapping)
		assert.Equal(t, mapping.Name, "other")
	})

	t.Run("Finds by name, returning nil if not exist", func(t *testing.T) {
		mapping := m.FindByName("nope")

		assert.Nil(t, mapping)
	})
}

func mockMappings() Mappings {
	return Mappings{
		Mapping{
			Name:        "name",
			Aliases:     []string{"name", "n"},
		},
		Mapping{
			Name:        "other",
			Aliases:     []string{"other"},
		},
	}
}
