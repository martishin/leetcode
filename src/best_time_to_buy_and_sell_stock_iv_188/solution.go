package best_time_to_buy_and_sell_stock_iv_188

import "fmt"

func maxProfit(k int, prices []int) int {
	n := len(prices)

	hold := make([][]int, k)
	notHold := make([][]int, k)

	for i := range k {
		hold[i] = make([]int, n)
		notHold[i] = make([]int, n)
		hold[i][0] = -prices[0]
	}

	for i := 0; i < k; i++ {
		for j := 1; j < n; j++ {
			if i == 0 {
				hold[i][j] = max(hold[i][j-1], -prices[j])
			} else {
				hold[i][j] = max(hold[i][j-1], notHold[i-1][j-1]-prices[j])
			}
			notHold[i][j] = max(hold[i][j]+prices[j], notHold[i][j-1])
		}
	}

	return notHold[k-1][n-1]
}

func Test() {
	fmt.Println(maxProfit(2, []int{2, 4, 1}))
	fmt.Println(maxProfit(2, []int{3, 2, 6, 5, 0, 3}))
}
