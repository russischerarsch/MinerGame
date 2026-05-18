package miners

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"

	"github.com/google/uuid"
)

const (
	LittleTypeMiner  string        = "littleMiner"
	LittleSalary     int           = 5
	LittleAbillity   int           = 1
	LittleRestNeeded time.Duration = 3
)

type LittleMiner struct {
	ID          uuid.UUID
	Salary      int
	TypeMiner   string
	Possibility *atomic.Int64
	Abillity    int
	RestNeeded  time.Duration
}

func (l *LittleMiner) Run(ctx context.Context) <-chan int64 {
	transferPoint := make(chan int64)
	go func() {
		defer func() {
			close(transferPoint)
			fmt.Printf("Im a miner number %v. I finished my work ", l.ID)
		}()
		for l.Possibility.Load() > 0 {
			select {

			case <-ctx.Done():
				fmt.Printf("Im a little miner number %v. I abruptly stopped my work ", l.ID)
				return
			case <-time.After(l.RestNeeded):
				transferPoint <- int64(l.Abillity)
				l.Possibility.Add(-1)
			}
		}
	}()
	return transferPoint
}

func NewLittleMiner() *LittleMiner {
	possibility := &atomic.Int64{}
	possibility.Add(30)
	return &LittleMiner{
		ID:          uuid.New(),
		Salary:      LittleSalary,
		TypeMiner:   LittleTypeMiner,
		Possibility: possibility,
		Abillity:    LittleAbillity,
		RestNeeded:  LittleRestNeeded,
	}
}

func (l *LittleMiner) Info() MinerInfo {
	return MinerInfo{
		ID:        l.ID,
		MinerType: LittleTypeMiner,
		PossLeft:  l.Possibility.Load(),
	}
}
