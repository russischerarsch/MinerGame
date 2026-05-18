package miners

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"

	"github.com/google/uuid"
)

const (
	MiddleTypeMiner  string        = "MiddleMiner"
	MiddleSalary     int           = 50
	MiddleAbillity   int           = 3
	MiddleRestNeeded time.Duration = 2
)

type MiddleMiner struct {
	ID          uuid.UUID
	Salary      int
	TypeMiner   string
	Possibility *atomic.Int64
	Abillity    int
	RestNeeded  time.Duration
}

func (l *MiddleMiner) Run(ctx context.Context) <-chan int64 {
	transferPoint := make(chan int64)
	go func() {
		defer func() {
			close(transferPoint)
			fmt.Printf("Im a miner number %v. I finished my work ", l.ID)
		}()
		for l.Possibility.Load() > 0 {
			select {

			case <-ctx.Done():
				fmt.Printf("Im a middle miner number %v. I abruptly stopped my work ", l.ID)
				return
			case <-time.After(l.RestNeeded):
				transferPoint <- int64(l.Abillity)
				l.Possibility.Add(-1)
			}
		}
	}()
	return transferPoint
}

func NewMiddleMiner() *MiddleMiner {
	possibility := &atomic.Int64{}
	possibility.Add(45)
	return &MiddleMiner{
		ID:          uuid.New(),
		Salary:      MiddleSalary,
		TypeMiner:   MiddleTypeMiner,
		Possibility: possibility,
		Abillity:    MiddleAbillity,
		RestNeeded:  MiddleRestNeeded,
	}
}
func (m MiddleMiner) Info() MinerInfo {
	return MinerInfo{
		ID:        m.ID,
		MinerType: MiddleTypeMiner,
		PossLeft:  m.Possibility.Load(),
	}
}
