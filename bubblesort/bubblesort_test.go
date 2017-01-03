package main

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	cases := []struct {
		in, want []int
	}{
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{5, 5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5, 5}},
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{0, 1, 2, 3, 4, 3, 2, 1, 0}, []int{0, 0, 1, 1, 2, 2, 3, 3, 4}},
	}
	for _, c := range cases {
		//got := Bubblesort(c.in)
		got := make([]int, len(c.in))
		copy(got, c.in)
		Bubblesort(got)
		//if got != c.want {
		if !reflect.DeepEqual(got, c.want) {

			t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
