package main

import(
	
)

func BestTimeToBuyAndSellStocks(prices []int) int{
	minPrice := prices[0]
	maxProfit := 0

	for _, price := range prices{
		profit := price - minPrice
		if profit > maxProfit {
			maxProfit = profit
		}
		if price < minPrice{
			minPrice = price
		}
	}
	return maxProfit
}