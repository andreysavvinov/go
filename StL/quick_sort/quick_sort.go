package quick_sort

func QuickSort[T any](arr []T, less func(a, b T) bool) {
	if len(arr) < 2 {
		return
	}

	quickSort(arr, 0, len(arr)-1, less)
}

func quickSort[T any](arr []T, left, right int, less func(a, b T) bool) {
	if left >= right {
		return
	}

	pivotIndex := partition(arr, left, right, less)

	quickSort(arr, left, pivotIndex-1, less)
	quickSort(arr, pivotIndex+1, right, less)
}

func partition[T any](
	arr []T,
	left, right int,
	less func(a, b T) bool,
) int {
	pivot := arr[right]

	i := left

	for j := left; j < right; j++ {
		if less(arr[j], pivot) {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	arr[i], arr[right] = arr[right], arr[i]

	return i
}
