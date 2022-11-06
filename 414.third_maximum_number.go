import "math"

func thirdMax(nums []int) int {
	b := &BigThree{
		max:    math.MinInt64,
		second: math.MinInt64,
		third:  math.MinInt64,
	}

	for _, num := range nums {
		b.Set(num)
	}
	return b.Result()
}

type BigThree struct {
	max    int
	second int
	third  int
}

func (b *BigThree) Set(num int) {
	if num == b.max || num == b.second || num == b.third {
		return
	}
	if num > b.max {
		b.max, b.second, b.third = num, b.max, b.second
		return
	}
	if num > b.second {
		b.second, b.third = num, b.second
		return
	}
	if num > b.third {
		b.third = num
	}
}

func (b *BigThree) Result() int {
	if b.third == math.MinInt64 {
		return b.max
	}
	return b.third
}

