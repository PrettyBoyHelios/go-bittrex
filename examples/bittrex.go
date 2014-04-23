package main

import (
	"fmt"
	"github.com/toorop/go-bittrex"
)

const (
	API_KEY = ""
)

func main() {
	// Bittrex client
	bittrex := bittrex.New(API_KEY)

	// Get markets
	//markets, err := bittrex.GetMarkets()
	//fmt.Println(err, markets)

	// Get Ticker (BTC-VTC)
	//ticker, err := bittrex.GetTicker("BTC-VTC")
	//fmt.Println(err, ticker)

	// Get market summaries
	//marketSummaries, err := bittrex.GetMarketSummaries()
	//fmt.Println(err, marketSummaries)

	// Get orders book
	/*orderBook, err := bittrex.GetOrderBook("BTC-QBC", "both", 100)
	fmt.Println(err, orderBook)
	*/

	// Market historu
	marketHistory, err := bittrex.GetMarketHistory("BTC-QBC", 100)
	fmt.Println(err, marketHistory)
}