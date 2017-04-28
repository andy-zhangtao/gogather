package zsort

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickSort(t *testing.T) {
	data := sort.IntSlice{12, 45, 89, 01, -2, -99, 28}
	dest := [...]int{-99, -2, 01, 12, 28, 45, 89}
	QuickSort(data)

	for i, d := range data {
		assert.Equal(t, d, dest[i], "They should be equal")
	}
}

func Benchmark_quickSort(b *testing.B) {
	data := sort.IntSlice{12, 45, 89, 01, -2, -99, 28}
	b.StartTimer()
	QuickSort(data)
	b.StopTimer()
}

func ExampleQuickSort() {
	// declare a array
	// this array must implenet sort.Inerface
	data := sort.IntSlice{22, 34, 3, 40, 18, 4}
	QuickSort(data)
}
