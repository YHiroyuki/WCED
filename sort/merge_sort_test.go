package sort

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func sampleData(n int) ([]int, []int) {
	expected := []int{}
	for i := 0; i < 19; i++ {
		expected = append(expected, i)
	}

	sample := make([]int, len(expected))
	copy(sample, expected)

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(sample), func(i, j int) { sample[i], sample[j] = sample[j], sample[i] })

	return expected, sample
}

func TestQuickSort(t *testing.T) {
	expected, sample := sampleData(1000)

	actual := QuickSort(sample)

	if !reflect.DeepEqual(expected, actual) {
		t.Fatal()
	}
}

func TestMergeSort(t *testing.T) {
	expected, sample := sampleData(19)

	actual := MergeSort(sample)

	if !reflect.DeepEqual(expected, actual) {
		t.Fatal()
	}
}

func TestQuickSort100(t *testing.T) {
	_, sample := sampleData(100)

	start := time.Now()
	QuickSort(sample)
	end := time.Now()
	t.Logf("%f秒\n", (end.Sub(start)).Seconds())
}

func TestB(t *testing.T) {
	_, sample := sampleData(100000000)

	start := time.Now()
	QuickSort(sample)
	end := time.Now()
	t.Logf("%f秒\n", (end.Sub(start)).Seconds())
}

func TestMergeSort100(t *testing.T) {
	_, sample := sampleData(100)

	start := time.Now()
	MergeSort(sample)
	end := time.Now()
	t.Logf("%f秒\n", (end.Sub(start)).Seconds())
}

func TestA(t *testing.T) {
	_, sample := sampleData(100000000)

	start := time.Now()
	MergeSort(sample)
	end := time.Now()
	t.Logf("%f秒\n", (end.Sub(start)).Seconds())
}
