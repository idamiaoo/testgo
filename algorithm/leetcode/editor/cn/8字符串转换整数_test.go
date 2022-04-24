package leetcode

import (
	"fmt"
	"math"
	"testing"
)

func myAtoi(s string) int {
	var ans int
	var start int

	for start = 0; start < len(s); start++ {
		fmt.Println(s[start])
		if s[start] != ' ' {
			break
		}
	}
	var sign = 1
	for i := start; i < len(s); i++ {
		if (s[i] == '+' || s[i] == '-') && i == start {
			if s[i] == '-' {
				sign = -1
			}
		} else if s[i] >= '0' && s[i] <= '9' {
			num := int(int(s[i] - '0'))
			if ans > math.MaxInt32/10 || (ans == math.MaxInt32/10 && num > math.MaxInt32%10) {
				return math.MaxInt32
			}
			if ans < math.MinInt32/10 || (ans == math.MinInt32/10 && -num < math.MinInt32%10) {
				return math.MinInt32
			}
			ans = ans*10 + (sign * num)
		} else {
			break
		}
	}
	return ans
}

func TestMyAtoi(t *testing.T) {
	t.Log(myAtoi("   -42"))
}

/*
class Solution {
    public int myAtoi(String s) {
        int sign = 1;
        int res = 0;
        int m = s.length();
        int i = 0;
        while(i < m && s.charAt(i)==' '){
            i++;
        }
        int start = i;
        for(; i < m; i++){
            char c = s.charAt(i);
            if(i==start && c=='+'){
                sign = 1;
            }else if(i==start && c=='-'){
                sign = -1;
            }else if(Character.isDigit(c)){
                int num = c-'0';
                if(res > Integer.MAX_VALUE/10 || (res == Integer.MAX_VALUE/10&&num>Integer.MAX_VALUE%10)){
                    return Integer.MAX_VALUE;
                }

                if(res < Integer.MIN_VALUE/10 || (res == Integer.MIN_VALUE/10&&-num<Integer.MIN_VALUE%10)){
                    return Integer.MIN_VALUE;
                }
                res = res*10+sign*num;
            }else{
                break;
            }
        }
        return res;
    }
}
*/
