package aoc15

func MinOf(vars ...int) int {
	if len(vars) == 0 {
		return -1
	}
	min := vars[0]
	for _, i := range vars {
		if min > i {
			min = i
		}
	}
	return min
}

func MaxOf(vars ...int) int {
	max := vars[0]
	for _, i := range vars {
		if max < i {
			max = i
		}
	}
	return max
}

func Sum(vars ...int) int {
	sum := 0
	for _, i := range vars {
		sum += i
	}
	return sum
}

func Map[T, V any](ts []T, fn func(T) V) []V {
	result := make([]V, len(ts))
	for i, t := range ts {
		result[i] = fn(t)
	}
	return result
}
