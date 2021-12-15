package leetcode

import (
	"testing"
)

func singleNumber(nums []int) []int {
	var x int
	for _, v := range nums {
		x ^= v
	}
	l := x & (-x)
	var n1, n2 int
	for _, v := range nums {
		if l&v > 0 {
			n1 ^= v
		} else {
			n2 ^= v
		}
	}
	return []int{n1, n2}
}

func TestSingleNumber(t *testing.T) {
	t.Log(singleNumber([]int{1, 2, 1, 3, 2, 5}))
}

// 0000 0101
// 0000 0011
// 0000 0110
// 1111 1010
/*
func TestSearch(t *testing.T) {
	v := []string{
		"/v1/sys/organization/:id::DELETE", "/v1/sys/organization/:id::PUT", "/v1/sys/organization::GET", "/v1/sys/organization::POST", "/v1/sys/user/:id::DELETE", "/v1/sys/user/:id::PUT", "/v1/sys/user::GET", "/v1/sys/user::POST", "/v1/sys/user_role/:id::DELETE", "/v1/sys/user_role/:id::PUT", "/v1/sys/user_role::GET", /v1/sys/user_role::POST AddWafRegionBlockRule],
	}
	t.Log("/v1/sys/organization/:id::DELETE"
	"/v1/sys/log_audit::GET")
}
*/
