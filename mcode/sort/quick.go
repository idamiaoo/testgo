package sort

func QuickSort(nums []int, l, r int) {

}

func TopK(nums []int, k int) {

}

func BubbleSort(nums []int, l, r int) {
	for j := r - 1; j > l; j-- {
		for i := l; i <= j; i++ {
			if nums[i] > nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
			}
		}
	}
}
