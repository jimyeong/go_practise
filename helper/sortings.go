package sortings

func BinarySearch(arr []int, target int) int {

	var end int = len(arr) - 1
	var start int = 0

	var index int = -1
	for end-start > 1 {
		mid := (start + end) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			start = mid + 1
		} else if arr[mid] > target {
			end = mid - 1
		}
	}
	return index
}
