package slice

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestAppend(t *testing.T) {
	Append()
}

func TestCopy(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6}
	b := []int{3, 2, 1}
	printSlice(a)
	printSlice(b)
	copy(a[3:], b)
	printSlice(a)
}

func TestRange(t *testing.T) {
	i1, i2, i3, i4, i5 := 1, 2, 3, 4, 5
	a := []*int{&i1, &i2, &i3, &i4, &i5}
	for _, p := range a {
		*p = 1
	}
	for _, t := range a {
		fmt.Println(*t)
	}
}

type SS struct {
	ss []int
}

func TestCopyV2(t *testing.T) {
	ts := &SS{
		ss: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
	}

	go func() {
		for {
			log.Println(ts.ss)
			<-time.After(1 * time.Second)
		}
	}()

	its := &ts.ss

	for {
		if len(*its) == 0 {
			return
		}
		<-time.After(1 * time.Second)
		*its = (*its)[:len(*its)-1]
		log.Println("its ", its)
	}
}

func change(s []int) {
	for i := range s {
		s[i] = 1
	}
}

func TestCopyV3(t *testing.T) {
	ss := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	change(ss)
	log.Println(ss)
}

func TestMapS(t *testing.T) {
	a := make(map[string][]string)
	a["bohler"] = append(a["bohler"], "sdsdf")
	a["bohler"] = append(a["bohler"], "dsffffff")
	fmt.Println(a["bohler"])
}

func TestRangeAndAppend(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6}

	for i, v := range a {
		fmt.Println("i:", i)
		if v == 1 {
			a = a[3:]
		}
		fmt.Println(v, a[i])
	}
	fmt.Println(a)
}

type Grade struct {
	Name  string
	Score int
}

func TestRangeModify(t *testing.T) {
	grades := []Grade{
		{
			Name:  "carl",
			Score: 100,
		},
		{
			Name:  "tyler",
			Score: 100,
		},
	}

	for _, v := range grades {
		if v.Name == "tyler" {
			t.Log("modify")
			v.Score = 0
		}
	}
	t.Logf("%#v", grades)
}
