package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/YHiroyuki/WCED/sort"
)

func main() {

	sample := []int{}
	for i := 0; i < 20; i++ {
		sample = append(sample, i)
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(sample), func(i, j int) { sample[i], sample[j] = sample[j], sample[i] })
	fmt.Println(sample)

	ans := sort.MergeSort(sample)

	fmt.Println(ans)
}
