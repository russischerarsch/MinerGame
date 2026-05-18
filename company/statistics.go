package company

import (
	"sync/atomic"
)

type CompanyStatistics struct {
	AllMiners   map[string]int
	TotalEarned *atomic.Int64
	Balance     *atomic.Int64
}

func NewCompanyStatistics() *CompanyStatistics {
	return &CompanyStatistics{
		AllMiners:   make(map[string]int),
		TotalEarned: &atomic.Int64{},
		Balance:     &atomic.Int64{},
	}
}
func (c CompanyStatistics) GetBalance() int {
	return int(c.Balance.Load())
}
func (c CompanyStatistics) GetTotalEarned() int {
	return int(c.TotalEarned.Load())
}
func (c CompanyStatistics) GetAllMiners() map[string]int {
	tmp := make(map[string]int)
	for k, v := range c.AllMiners {
		tmp[k] = v
	}
	return tmp
}
