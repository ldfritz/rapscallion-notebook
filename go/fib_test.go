package fib

import (
	"fmt"
	"testing"
)

var testMaxes = []struct {
	arg      int
	expected int
}{
	{0, 0},
	{1, 1},
	{2, 2},
	{3, 3},
	{4, 3},
	{5, 5},
	{6, 5},
	{7, 5},
	{8, 8},
	{9, 8},
}

var testSets = []struct {
	arg      int
	expected []int
}{
	{0, []int{0}},
	{1, []int{0, 1, 1}},
	{2, []int{0, 1, 1, 2}},
	{3, []int{0, 1, 1, 2, 3}},
	{4, []int{0, 1, 1, 2, 3}},
	{5, []int{0, 1, 1, 2, 3, 5}},
	{6, []int{0, 1, 1, 2, 3, 5}},
	{7, []int{0, 1, 1, 2, 3, 5}},
	{8, []int{0, 1, 1, 2, 3, 5, 8}},
	{9, []int{0, 1, 1, 2, 3, 5, 8}},
}

func TestFibMax1(t *testing.T) {
	for _, v := range testMaxes {
		if actual := fibMax1(v.arg); actual != v.expected {
			t.Errorf("fibMax1(%d):\n  Expected: %d\n  Actual:   %d", v.arg, v.expected, actual)
		}
	}
}

func TestFibSet1(t *testing.T) {
	for _, pair := range testSets {
		actual := fibSet1(pair.arg)
		match := true
		if len(pair.expected) != len(actual) {
			match = false
		}
		for i, v := range actual {
			if v != pair.expected[i] {
				match = false
			}
		}
		if !match {
			t.Errorf("fibSet1(%d):\n  Expected: %d\n  Actual:   %d", pair.arg, pair.expected, actual)
		}
	}
}

func TestIncAndDecFib(t *testing.T) {
	a := 2
	b := 3
	args := fmt.Sprintf("%d, %d", a, b)
	incFib(&a, &b)
	if a != 3 && b != 5 {
		t.Errorf("incFib(%s):\n  Expected: 3 5 (side-effect)\n  Actual:   %d %d", args, a, b)
	}
	args = fmt.Sprintf("%d, %d", a, b)
	incFib(&a, &b)
	if a != 5 && b != 8 {
		t.Errorf("incFib(%s):\n  Expected: 5 8 (side-effect)\n  Actual:   %d %d", args, a, b)
	}
	args = fmt.Sprintf("%d, %d", a, b)
	decFib(&a, &b)
	if a != 3 && b != 5 {
		t.Errorf("incFib(%s):\n  Expected: 3 5 (side-effect)\n  Actual:   %d %d", args, a, b)
	}
	args = fmt.Sprintf("%d, %d", a, b)
	decFib(&a, &b)
	if a != 2 && b != 3 {
		t.Errorf("incFib(%s):\n  Expected: 5 8 (side-effect)\n  Actual:   %d %d", args, a, b)
	}
}

/*
func TestFibSet2(t *testing.T) {
	for _, pair := range testSets {
		actual := fibSet2(pair.arg)
		match := true
		if len(pair.expected) != len(actual) {
			match = false
		}
		for i, v := range actual {
			if v != pair.expected[i] {
				match = false
			}
		}
		if !match {
			t.Errorf("fibSet2(%d):\n  Expected: %d\n  Actual:   %d", pair.arg, pair.expected, actual)
		}
	}
}
*/
