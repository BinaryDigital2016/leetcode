package others

func moveZeroes(nums []int) {
	flag := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[flag] = nums[flag], nums[i]
			flag++
		}
	}
}
