package sort

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func TestMergeSort(t *testing.T) {
	expected := []int{}
	for i := 0; i < 19; i++ {
		expected = append(expected, i)
	}

	sample := make([]int, len(expected))
	copy(sample, expected)

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(sample), func(i, j int) { sample[i], sample[j] = sample[j], sample[i] })

	actual := MergeSort(sample)
	t.Log(sample)
	t.Log(expected)

	t.Log(actual)
	if !reflect.DeepEqual(sample, actual) {
		t.Fatal()
	}
}
