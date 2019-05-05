package kata02

func chop(key int, array []int, offset, length int) int {
	remaining := length - offset
	if remaining == 0 {
		return -1
	}

	point := offset + remaining/2
	if array[point] == key {
		return point
	}
	if array[point] > key {
		return chop(key, array, offset, point)
	}
	return chop(key, array, point+1, length)
}

func Chop(key int, array []int) int {
	return chop(key, array, 0, len(array))
}
