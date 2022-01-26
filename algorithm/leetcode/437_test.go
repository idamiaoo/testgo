package leetcode

func pathSum(root *TreeNode, targetSum int) int {
	sums := make(map[int]int)
	sums[0] = 1
	var ans int
	var dfs func(*TreeNode, int)
	dfs = func(root *TreeNode, cur int) {
		if root == nil {
			return
		}
		cur += root.Val
		ans += sums[cur-targetSum]
		sums[cur]++
		dfs(root.Left, cur)
		dfs(root.Right, cur)
		sums[cur]--
	}
	dfs(root, 0)
	return ans
}
