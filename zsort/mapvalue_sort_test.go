package zsort

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSortByValue(t *testing.T) {

	data := map[string]int{
		"abc": 9,
		"dbd": 2,
		"fff": 5,
		"efa": 4,
	}

	dest := []Pair{
		{"abc",9},
		{"fff",5},
		{"efa",4},
		{"dbd",2},
	}
	resp := SortByValue(data)

	for i, r := range resp{
		assert.Equal(t, r, dest[i], "They should be equal")
	}
}
