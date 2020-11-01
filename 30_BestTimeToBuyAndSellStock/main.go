package main

func maxProfit(prices []int) int {
	var minCost int = 1 << 31 - 1
	maxProfit := 0
	for _, val := range prices {
		if minCost > val {
			minCost = val
		} else if maxProfit < val - minCost {
			maxProfit = val - minCost
		}
	}
	return maxProfit
}

func main() {

}
