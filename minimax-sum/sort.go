package main

func quickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	pivotIndex := partitio(arr)
	quickSort(arr[:pivotIndex])
	quickSort(arr[pivotIndex+1:])
}

func partitio(arr []int) int {
	pivot := arr[0]
	left, right := 1, len(arr)-1

	for left <= right {
		for left <= right && arr[left] < pivot {
			left++
		}

		for left <= right && arr[right] > pivot {
			right--
		}

		if left <= right {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}
	}

	arr[0], arr[right] = arr[right], arr[0]
	return right
}
