package utils

func AddTwoNums(x, y int) (sum int) {
	sum = x + y
	return
}

func AddUp(x int) int {
	if x < 2 {
		return x
	}
	sum := 0
	for i := 0; i <= x; i++ {
		sum += i
	}
	return sum
}

func AddUpMore(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

func AddThreeNums(x, y, z int) (sum int) {
	sum = x + y + z
	return
}
