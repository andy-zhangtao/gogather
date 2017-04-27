package zsort

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {
	data := sort.IntSlice{22, 34, 3, 40, 18, 4}

	dest := [...]int{40, 34, 22, 18, 4, 3}
	BubbleSort(data)

	for i, d := range data {
		assert.Equal(t, d, dest[i], "They should be equal")
	}
}

func Benchmark_bubbleSort(b *testing.B) {
	data := sort.IntSlice{22, 34, 3, 40, 18, 4, 22, 34, 3, 40, 18, 4, 22, 34, 3, 40, 18, 4}
	b.StartTimer()
	BubbleSort(data)
}

func ExampleBubbleSort() {
	// declare a array
	// this array must implenet sort.Inerface
	data := sort.IntSlice{22, 34, 3, 40, 18, 4}
	BubbleSort(data)
}
