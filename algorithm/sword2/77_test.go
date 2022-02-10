package sword2

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	h1, h2 := head, splitNode(head)
	h1 = sortList(h1)
	h2 = sortList(h2)
	return mergeNode(h1, h2)
}

func mergeNode(l1, l2 *ListNode) *ListNode {
	var head = &ListNode{}
	temp := head
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			temp.Next = l1
			l1 = l1.Next
		} else {
			temp.Next = l2
			l2 = l2.Next
		}
		temp = temp.Next
	}
	if l1 != nil {
		temp.Next = l1
	}
	if l2 != nil {
		temp.Next = l2
	}
	return head.Next
}

// 1 2 3 4 5 6 7
func splitNode(head *ListNode) *ListNode {
	var slow, fast = head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	temp := slow.Next
	slow.Next = nil
	return temp
}

func mergeSort(nums []int, left, right int) []int {
	if len(nums) < 2 {
		return nums
	}
	mid := (left + right) >> 1
	a := mergeSort(nums, left, mid)
	b := mergeSort(nums, mid+1, right)
	return merge77(a, b)
}

func merge77(nums1, nums2 []int) []int {
	var i, j int
	temp := make([]int, 0, len(nums1)+len(nums2))
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			temp = append(temp, nums1[i])
			i++
		} else {
			temp = append(temp, nums2[j])
			j++
		}
	}
	for i < len(nums1) {
		temp = append(temp, nums1[i])
		i++
	}
	for j < len(nums2) {
		temp = append(temp, nums2[j])
		j++
	}
	return temp
}

/*
1 3 4 7 9

2 5 6 8 10
*/
