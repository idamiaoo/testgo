package tencent2_test

import "testing"

/*
编程题：
1、写出Linux shell能完成如下工作的命令
(1) 统计一个文本的行数
(2) 修改文件权限的命令
(3) 有1个log文件如下格式（qq号|时间|行为）：
123456|2019-03-15 08:55:52|login
12345688|2019-03-15 08:58:27|login
12378900|2019-03-15 10:23:03|logout
8899000|2019-03-15 12:01:31|login
123456|2019-03-15 14:31:42|logout
...
请用shell脚本统计该log文件号码为123456的login次数

2、实现一个函数，判断整数是否为2的整幂
// Golang
func isPower2(n int) bool {

}

3、实现快速排序算法
// Golang
func qsort(arr []int, start, end int) {

}
*/

/*
1.
 (1) find
 (2) chmod
 (3)

*/

// 2
func isPower2(n int) bool {
	if n == 1 {
		return true
	}
	return findPower2(1, n, n)
}

func findPower2(start, end, n int) bool {
	if end-start == 1 {
		return false
	}

	x := (start + end) / 2
	if x*x == n {
		return true
	} else if x*x < n {
		return findPower2(x, end, n)
	} else {
		return findPower2(start, x, n)
	}
}

func TestIsPower2(t *testing.T) {
	t.Log(isPower2(16))
}

// 3
func qsort(arr []int, start, end int) {
	if start < end {
		p := parti(arr, start, end)
		qsort(arr, start, p-1)
		qsort(arr, p+1, end)
	}
}

func parti(arr []int, start, end int) int {
	b := arr[end]
	i := start - 1

	for j := 0; j < end-1; j++ {
		if arr[j] <= b {
			i++
			arr[j], arr[i] = arr[i], arr[j]
		}
	}
	arr[i+1], arr[end] = arr[end], arr[i+1]
	return i + 1
}
