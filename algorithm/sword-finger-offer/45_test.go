package sword_finger_offer

func minNumber(nums []int) string {
	return ""
}

func LSort(array []string, left, right int) {
	if left < right {

	}
}

func LPartition(arr []string, left, right int) {
	p := arr[right]

	n := left - 1

	for j := 0; j < right; j++ {
		if arr[j]+p < p+arr[j] {
			n++
			arr[n], arr[j] = arr[j], arr[n]
		}
	}
}
