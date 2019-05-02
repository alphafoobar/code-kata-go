package kata02

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func calculateMid(remaining int) int {
	return min(remaining-1, remaining/2)
}

func foundIndex(key int, array []int, offset int) int {
	if array[offset] == key {
		return offset
	}
	return -1
}

func chop(key int, array []int, offset, length int) int {
	remaining := length - offset

	if remaining == 0 {
		return -1
	}
	if remaining == 1 {
		return foundIndex(key, array, offset)
	}

	midPoint := calculateMid(remaining)

	point := offset + midPoint
	if array[point] == key {
		return point
	}
	if array[point] > key {
		return chop(key, array, offset, point)
	}
	return chop(key, array, point, length)
}

func Chop(key int, array []int) int {
	return chop(key, array, 0, len(array))
}
