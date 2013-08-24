package fanalysis

import (
	"github.com/themoriarty/yfinance"
	"time"
)

type Event struct{
	Date time.Time
}

func FindEvents(prices []yfinance.Price, algo func(yfinance.Price, []yfinance.Price) bool) (ret []Event){
	ret = nil
	for i, p := range(prices){
		if algo(p, prices[:i]){
			ret = append(ret, Event{p.Date})
		}
	}
	return ret
}