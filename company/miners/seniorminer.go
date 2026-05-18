package miners

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/google/uuid"
)

const (
	LeadTypeMiner  string        = "LeadMiner"
	LeadSalary     int           = 450
	LeadAbillity   int           = 10
	LeadRestNeeded time.Duration = 1
)

type LeadMiner struct {
	ID          uuid.UUID
	Salary      int
	TypeMiner   string
	Possibility *atomic.Int64
	Abillity    int
	RestNeeded  time.Duration
	mtx         sync.Mutex
}

func (l *LeadMiner) Run(ctx context.Context) <-chan int64 {
	transferPoint := make(chan int64)
	go func() {
		defer func() {
			close(transferPoint)
			fmt.Printf("Im a miner number %v. I finished my work ", l.ID)
		}()
		for l.Possibility.Load() > 0 {
			select {

			case <-ctx.Done():
				fmt.Printf("Im a lead miner number %v. I abruptly stopped my work ", l.ID)
				return
			case <-time.After(l.RestNeeded):
				transferPoint <- int64(l.Abillity)
				l.Possibility.Add(-1)
				l.mtx.Lock()
				l.Abillity += 3
				l.mtx.Unlock()
			}
		}
	}()
	return transferPoint
}

func NewLeadMiner() *LeadMiner {
	possibility := &atomic.Int64{}
	possibility.Add(60)
	return &LeadMiner{
		ID:          uuid.New(),
		Salary:      LeadSalary,
		TypeMiner:   LeadTypeMiner,
		Possibility: possibility,
		Abillity:    LeadAbillity,
		RestNeeded:  LeadRestNeeded,
	}
}
func (l LeadMiner) Info() MinerInfo {
	return MinerInfo{
		ID:        l.ID,
		MinerType: LeadTypeMiner,
		PossLeft:  l.Possibility.Load(),
	}
}
