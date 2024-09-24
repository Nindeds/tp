package poulpe

func Ftmissing(nums []int) int {
	nombre := make(map[int]bool)
	for _, num := range nums {
		nombre[num] = true
	}
	for i := 0; i <= 9 && i >= 0; i++ {
		if !nombre[i] {
			return i
		}
	}
	return 0
}
