package array

func buildArray(nums []int) []int {
	ans := make([]int, len(nums))

	for index, num := range nums {
		ans[index] = nums[num]
	}

	return ans
}
