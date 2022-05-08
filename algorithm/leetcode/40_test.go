package leetcode

import (
	"sort"
	"testing"
)

//给定一个候选人编号的集合 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
//
// candidates 中的每个数字在每个组合中只能使用 一次 。
//
// 注意：解集不能包含重复的组合。
//
//
//
// 示例 1:
//
//
//输入: candidates = [10,1,2,7,6,1,5], target = 8,
//输出:
//[
//[1,1,6],
//[1,2,5],
//[1,7],
//[2,6]
//]
//
// 示例 2:
//
//
//输入: candidates = [2,5,2,1,2], target = 5,
//输出:
//[
//[1,2,2],
//[5]
//]
//
//
//
// 提示:
//
//
// 1 <= candidates.length <= 100
// 1 <= candidates[i] <= 50
// 1 <= target <= 30
//
// Related Topics 数组 回溯 👍 930 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var (
		ans          [][]int
		path         []int
		sum          int
		backtracking func(index int)
	)
	backtracking = func(index int) {
		if sum == target {
			ans = append(ans, append([]int{}, path...))
			return
		}
		if sum > target {
			return
		}
		for i := index; i < len(candidates); i++ {
			if i > index && candidates[i] == candidates[i-1] {
				continue
			}
			sum += candidates[i]
			path = append(path, candidates[i])
			backtracking(i + 1)
			sum -= candidates[i]
			path = path[:len(path)-1]
		}
	}
	backtracking(0)
	return ans
}

func TestCombinationSum2(t *testing.T) {
	var candidates = []int{10, 1, 2, 7, 6, 1, 5}
	var target = 8
	t.Log(combinationSum2(candidates, target))
}

//leetcode submit region end(Prohibit modification and deletion)
