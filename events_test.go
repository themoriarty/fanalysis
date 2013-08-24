package fanalysis

import (
	"github.com/themoriarty/yfinance"
	"testing"
)

func TestOnlyOldData(t* testing.T){
	yf := yfinance.Interface{}
	allPrices, _ := yf.GetPrices([]string{"MSFT"}, yfinance.Date(2009, 1, 1), yfinance.Date(2009, 12, 31))
	prices, _ := allPrices.Prices("MSFT")
	ok := true
	callsCount := 0
	FindEvents(prices, func(today yfinance.Price, history []yfinance.Price) bool{
		callsCount++
		for _, p := range(history){
			if !today.Date.After(p.Date){
				ok = false
				t.Error("got non-historical price", p, "for today: ", today)
			}
		}
		return false
	});
	if !ok || callsCount != len(prices){
		t.Error("check failed, calls count: ", callsCount)
	}
}

func TestFilteringFilters(t* testing.T){
	yf := yfinance.Interface{}
	allPrices, _ := yf.GetPrices([]string{"MSFT"}, yfinance.Date(2009, 1, 1), yfinance.Date(2009, 12, 31))
	prices, _ := allPrices.Prices("MSFT")
	callsCount := 0
	events := FindEvents(prices, func(today yfinance.Price, history []yfinance.Price) bool{
		defer func(){ callsCount++ }()
		if (callsCount % 2) == 0{
			return true
		}
		return false
	});
	for i, p := range(prices){
		if (i % 2) == 0 {
			if events[i / 2].Date != p.Date{
				t.Error("found invalid event", events[i / 2], "expected: ", p.Date)
			}
		}
	}
}