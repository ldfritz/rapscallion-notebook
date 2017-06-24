package fib

func fibMax1(max int) int {
	a, b := 0, 1
	for b <= max {
		a, b = b, a+b
	}
	return a
}

func fibSet1(max int) []int {
	var nums []int
	a, b := 0, 1
	for b <= max {
		nums = append(nums, a)
		a, b = b, a+b
	}
	nums = append(nums, a)
	return nums
}

func fibMax2(max int) (num int) {
	for num, b := 0, 1; b <= max; incFib(&num, &b) {
	}
	return
}

func incFib(a *int, b *int) {
	next := *a + *b
	*a = *b
	*b = next
}

func decFib(a *int, b *int) {
	prev := *b - *a
	*b = *a
	*a = prev
}
