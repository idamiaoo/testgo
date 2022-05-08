package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

func main() {
	var (
		N      int
		sTimes []string
	)
	fmt.Scan(&N)
	for i := 0; i < N; i++ {
		var t string
		n, _ := fmt.Scan(&t)
		if n == 0 {
			break
		}
		sTimes = append(sTimes, t)
	}
	sort(sTimes, 0, len(sTimes)-1)
	for _, v := range sTimes {
		fmt.Println(v)
	}
}

type Time struct {
	hour, minute, second, mSecond int
}

func ParseTime(s string) Time {
	vs := strings.Split(s, ":")
	vss := strings.Split(vs[2], ".")
	t := Time{}
	t.hour, _ = strconv.Atoi(vs[0])
	t.minute, _ = strconv.Atoi(vs[1])
	t.second, _ = strconv.Atoi(vss[0])
	t.mSecond, _ = strconv.Atoi(vss[1])
	return t
}

func (t Time) Less(t2 Time) bool {
	if t.hour < t2.hour {
		return true
	}
	if t.hour > t2.hour {
		return false
	}
	if t.hour == t2.hour {
		if t.minute < t2.minute {
			return true
		}
		if t.minute > t2.minute {
			return false
		}
		if t.minute == t2.minute {
			if t.second < t2.second {
				return true
			}
			if t.second > t2.second {
				return false
			}
			if t.second == t2.second {
				if t.mSecond < t2.mSecond {
					return true
				}
				if t.mSecond > t2.mSecond {
					return false
				}
				if t.mSecond == t2.mSecond {
					return false
				}
			}
		}
	}
	return false
}

func sort(times []string, left, right int) {
	if left < right {
		p := partition(times, left, right)
		sort(times, left, p-1)
		sort(times, p+1, right)
	}
}

func partition(times []string, left, right int) int {
	v := rand.Intn(right-left+1) + left
	times[v], times[right] = times[right], times[v]
	b := times[right]
	i := left - 1
	for j := left; j < right; j++ {
		if ParseTime(times[j]).Less(ParseTime(b)) {
			i++
			times[i], times[j] = times[j], times[i]
		}
	}
	if times[i+1] != times[right] {
		times[i+1], times[right] = times[right], times[i+1]
	}
	return i + 1
}
