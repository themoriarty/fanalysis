package fanalysis

import (
	"github.com/themoriarty/yfinance"
)

type History struct{
	Prices []yfinance.Price
}

func (this History) LastNDays(days int) History{
	return History{this.Prices[len(this.Prices) - days:]}
}

func (this History) Get(idx int) float64{
	return float64(this.Prices[idx].AdjustedClose)
}
func (this History) Len() int{
	return len(this.Prices)
}
func (this History) Yesterday() (yfinance.Price, bool){
	if len(this.Prices) > 0{
		return this.Prices[len(this.Prices) - 1], true
	}
	return yfinance.Price{}, false
}

