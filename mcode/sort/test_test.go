package sort

import (
	"fmt"
	"testing"
)

/*
1. 按快排的方式处理,但是找到 中间点后,比较 k 和 i+1 的大小
	k == i+1 return
	k < i+1  走快排左分支
	k > i+1 走快排右分支

	o (n)
*/

func topK(data []int, l, r, k int) {
	if l < r {
		fmt.Println(l, r)
		p := partition2(data, l, r)

		temp := p - l //数组前半部分元素个数

		if k == temp+1 {
			return
		} else if temp > k {
			topK(data, l, p-1, k)
		} else {
			topK(data, p+1, r, k-p)
		}

		/*
			if temp >= k {
				topK(data, l, p-1, k)
			} else {
				topK(data, p+1, r, k-p)
			}
		*/
	}
}

func quickSort(data []int, l, r int) {
	if l < r {
		p := partition(data, l, r)
		quickSort(data, l, p-1)
		quickSort(data, p+1, r)
	}
}

func partition(data []int, l, r int) int {
	p := data[r]
	i := l - 1
	for j := l; j < r; j++ {
		if data[j] <= p {
			i = i + 1
			data[j], data[i] = data[i], data[j]
		}
	}
	data[i+1], data[r] = data[r], data[i+1]
	return i + 1
}

func partition2(data []int, l, r int) int {
	p := data[r]
	i := l - 1
	for j := l; j < r; j++ {
		if data[j] >= p {
			i = i + 1
			data[j], data[i] = data[i], data[j]
		}
	}
	data[i+1], data[r] = data[r], data[i+1]
	return i + 1
}

func TestQuickSort(t *testing.T) {
	d := []int{44, 29, 4, 88, 2, 3, 8, 40, 68}
	quickSort(d, 0, len(d)-1)
	fmt.Println(d)
}

func TestTopK(t *testing.T) {
	d := []int{44, 29, 4, 88, 2, 3, 8, 40, 68, 9, 87, 78, 20, 97}
	topK(d, 0, len(d)-1, 5)
	fmt.Println(d)
}

/*
dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1] + prices[i])
              max(   选择 rest  ,           选择 sell      )

解释：今天我没有持有股票，有两种可能：
要么是我昨天就没有持有，然后今天选择 rest，所以我今天还是没有持有；
要么是我昨天持有股票，但是今天我 sell 了，所以我今天没有持有股票了。

dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0] - prices[i])
              max(   选择 rest  ,           选择 buy         )
*/

func TestTwo(t *testing.T) {
	a := make(chan int)
	b := make(chan int)
	go func() {
		for i := 1; i <= 100; i++ {
			<-a
			if i%2 == 1 {
				fmt.Println("g1: ", i)
			}
			b <- 1
		}
	}()

	go func() {
		for i := 1; i <= 100; i++ {
			<-b
			if i%2 == 0 {
				fmt.Println("g2: ", i)
			}
			a <- 1
		}
	}()

	a <- 1

	select {}
}

func TestBubbleSort(t *testing.T) {
	d := []int{44, 29, 4, 88, 2, 3, 8, 40, 68}
	BubbleSort(d, 0, len(d)-1)
	fmt.Println(d)
}
