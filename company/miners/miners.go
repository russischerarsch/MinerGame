package miners

import (
	"context"

	"github.com/google/uuid"
)

type Miner interface {
	Run(ctx context.Context) <-chan int64
	Info() MinerInfo
}

type MinerInfo struct {
	ID        uuid.UUID
	MinerType string
	PossLeft  int64 // сколько осталось сил у шахтера
}
