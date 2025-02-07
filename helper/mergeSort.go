package sortings

func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	l := MergeSort(arr[:mid])
	r := MergeSort(arr[mid:])

	return merge(l, r)
}

func merge(l, r []int) []int {
	result := make([]int, 0, len(l)+len(r))
	i, j := 0, 0

	// Merge elements in sorted order
	for i < len(l) && j < len(r) {
		if l[i] < r[j] {
			result = append(result, l[i])
			i++
		} else {
			result = append(result, r[j])
			j++
		}
	}
	result = append(result, l[i:]...)
	result = append(result, r[j:]...)
	return result
}
