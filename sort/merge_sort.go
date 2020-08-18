package sort

import (
	"math"
)

func MergeSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}

	center := int(math.Ceil(float64(length) / 2.0))
	beforeArr := MergeSort(arr[0:center])
	afterArr := MergeSort(arr[center:length])

	newArr := []int{}

	for {
		beforeLength := len(beforeArr)
		afterLength := len(afterArr)
		if beforeLength == 0 && afterLength == 0 {
			break
		} else if 0 < beforeLength && afterLength == 0 {
			newArr = append(newArr, beforeArr[0])
			beforeArr = beforeArr[1:]
		} else if beforeLength == 0 && 0 < afterLength {
			newArr = append(newArr, afterArr[0])
			afterArr = afterArr[1:]
		} else {
			if afterArr[0] < beforeArr[0] {
				newArr = append(newArr, afterArr[0])
				afterArr = afterArr[1:]
			} else {
				newArr = append(newArr, beforeArr[0])
				beforeArr = beforeArr[1:]
			}
		}
	}

	return newArr
}
