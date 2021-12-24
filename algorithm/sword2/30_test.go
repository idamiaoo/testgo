package sword2_test

import (
	"math/rand"
)

type RandomizedSet struct {
	m     map[int]int
	array []int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{
		m: make(map[int]int),
	}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.m[val]; ok {
		return false
	}
	this.array = append(this.array, val)
	this.m[val] = len(this.array) - 1
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.m[val]; !ok {
		return false
	}
	tail := this.array[len(this.array)-1]
	this.array[this.m[val]] = tail
	this.m[tail] = this.m[val]
	delete(this.m, val)
	this.array = this.array[:len(this.array)-1]
	return true
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	return this.array[rand.Intn(len(this.array))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
