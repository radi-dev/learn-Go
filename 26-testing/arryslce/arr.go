package arryslce

func Sum(arr [4]int) int {
	sum := 0
	for _, num := range arr {
		sum += num
	}
	return sum
}

func SumSlice(arr []int) int {
	sum := 0
	for _, num := range arr {
		sum += num
	}
	return sum
}

func SumAll(arrArr ...[]int) int {
	sum := 0
	for _, arr := range arrArr {
		sum += SumSlice(arr)
	}
	return sum
}

func SumAllTails(arrArr ...[]int) int {
	sum := 0
	for _, arr := range arrArr {
		sum += SumSlice(arr[1:])
	}
	return sum
}
