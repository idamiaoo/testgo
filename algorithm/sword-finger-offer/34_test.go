package sword_finger_offer

func pathSum(root *TreeNode, target int) [][]int {
	path := make([]int, 0, 1)
	data := make([][]int, 0, 1)
	backtrack(root, target, &path, &data)
	return data
}

func backtrack(root *TreeNode, target int, path *[]int, data *[][]int) {
	n := len(*path)

	*path = append(*path, root.Val)

	if root.Left == nil && root.Right == nil {
		var sum int
		for _, v := range *path {
			sum += v
		}
		if sum == target {
			temp := make([]int, len(*path))
			copy(temp, *path)
			*data = append(*data, temp)
		}
		return
	}

	if root.Left != nil {
		backtrack(root.Left, target, path, data)
	}

	if root.Right != nil {
		backtrack(root.Right, target, path, data)
	}
	*path = (*path)[:n]
}
