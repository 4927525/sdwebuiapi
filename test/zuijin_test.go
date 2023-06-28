package test

import (
	"fmt"
	"testing"
)

func TestZuijin(t *testing.T) {
	arr := []int{12, 16, 29, 34, 39, 43, 55, 64, 71, 89, 90, 9}
	zuijin := get_zuijin(40, arr)
	fmt.Println(zuijin)
}

func get_zuijin(this int, arr []int) int {
	min := 0
	if this == arr[0] {
		return arr[0]
	} else if this > arr[0] {
		min = this - arr[0]
	} else if this < arr[0] {
		min = arr[0] - this
	}

	for _, v := range arr {
		if v == this {
			return v
		} else if v > this {
			if min > v-this {
				min = v - this
			}
		} else if v < this {
			if min > this-v {
				min = this - v
			}
		}
	}

	for _, v := range arr {
		if this+min == v {
			return v
		} else if this-min == v {
			return v
		}
	}
	return min
}
