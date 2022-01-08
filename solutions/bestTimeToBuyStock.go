package solutions

/*

You are given an array prices where prices[i] is the price of a given stock on the ith day.

You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.

Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.
*/
func BestTimeToBuyStock(prices []int) int {
	left, right := 0, 1
	maxP := 0

	for right < len(prices) {
		// Compare left and right items in the prices arrays
		if prices[left] < prices[right] {
			// If price on the right day is greater than the price in the left day, get the profit
			profite := prices[right] - prices[left]
			// Return the max profite between maxP and profite
			if profite > maxP {
				maxP = profite
			}
		} else {
			// If the right price is less than the left, move forward the left pointer
			left = right
		}
		// Move forward on the right regardless
		right += 1
	}
	return maxP
}
