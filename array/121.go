package array

func maxProfit(prices []int) int {
	minStock := 100000
	maxDiff := 0

	for _, element := range prices {
		if minStock >= element {
			minStock = element
		}

		maxDiff = max(element-minStock, maxDiff)
	}

	return maxDiff
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
