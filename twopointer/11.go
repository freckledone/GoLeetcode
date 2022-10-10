package slidingwindow

func maxArea(heights []int) int {
	maxArea := 0
	i := 0
	j := len(heights) - 1
	for j >= i {
		currArea := (j - i) * min(heights[i], heights[j])
		if currArea > maxArea {
			maxArea = currArea
		}

		if heights[i] > heights[j] {
			j = j - 1
		} else {
			i = i + 1
		}
	}

	return maxArea
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a

}
