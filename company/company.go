package company

import (
	"context"
	"errors"
	"secPetProject/company/miners"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
)

type Company struct {
	AllMiners     map[string]map[uuid.UUID]miners.Miner
	IncomeChannel chan int64
	Statistic     *CompanyStatistics
	mtx           sync.Mutex
	CompanyCtx    context.Context
	CompanyStop   context.CancelFunc
	Equipment     Equipment
	TimeStart     time.Time
	TimeComplete  *time.Time
	GameTime      *time.Duration
}

func (c *Company) GetAllMiners() map[string]map[uuid.UUID]miners.Miner {
	tmp := make(map[string]map[uuid.UUID]miners.Miner)
	for types, minersMap := range c.AllMiners {
		tmp[types] = make(map[uuid.UUID]miners.Miner)
		for k, v := range minersMap {
			tmp[types][k] = v
		}
	}
	return tmp
}
func (c *Company) GetMinersByType(MinerType string) map[uuid.UUID]miners.Miner {
	tmp := make(map[uuid.UUID]miners.Miner)
	for minerType, MinersMap := range c.AllMiners {
		if MinerType == minerType {
			for k, v := range MinersMap {
				tmp[k] = v
			}
		}
	}
	return tmp
}

func (c *Company) BuyEquipment(equipment string) (Equipment, error) {
	equip := strings.ToLower(equipment)
	switch equip {
	case Pickaxe:
		if c.Statistic.Balance.Load() >= int64(PickaxesPrice) {
			c.Equipment.BuyPickaxes()
			c.Statistic.Balance.Add(-PickaxesPrice)
		} else {
			return Equipment{}, NotEnoughMoney
		}
	case Vents:
		if c.Statistic.Balance.Load() >= int64(VentsPrice) {
			c.Equipment.BuyVents()
			c.Statistic.Balance.Add(-VentsPrice)
		} else {
			return Equipment{}, NotEnoughMoney
		}
	case Trolleys:
		if c.Statistic.Balance.Load() >= int64(TrolleysPrice) {
			c.Equipment.BuyTrolleys()
			c.Statistic.Balance.Add(-TrolleysPrice)
		} else {
			return Equipment{}, NotEnoughMoney
		}
	}
	return c.Equipment, nil
}
func (c *Company) GetEquipment() Equipment {
	return c.Equipment
}
func (c *Company) Complete() error {
	if !c.Equipment.AllBought() {
		return errors.New("Not all equipment was bought")
	}
	tn := time.Now()
	c.TimeComplete = &tn
	overallTime := time.Since(c.TimeStart)
	c.GameTime = &overallTime
	return nil
}
func (c Company) GetStats() *CompanyStatistics {
	return c.Statistic
}
func (c *Company) HireMiner(MinerType string) (miners.Miner, error) {
	var miner miners.Miner
	switch MinerType {
	case miners.LittleTypeMiner:
		if c.Statistic.Balance.Load() >= int64(miners.LittleSalary) {
			miner = miners.NewLittleMiner()
			c.Statistic.Balance.Add(int64(-miners.LittleSalary))
		} else {
			return nil, NotEnoughMoney
		}
	case miners.MiddleTypeMiner:
		if c.Statistic.Balance.Load() >= int64(miners.MiddleSalary) {
			miner = miners.NewMiddleMiner()
			c.Statistic.Balance.Add(int64(-miners.MiddleSalary))
		} else {
			return nil, NotEnoughMoney
		}
	case miners.LeadTypeMiner:
		if c.Statistic.Balance.Load() >= int64(miners.LeadSalary) {
			miner = miners.NewLeadMiner()
			c.Statistic.Balance.Add(int64(-miners.MiddleSalary))
		} else {
			return nil, NotEnoughMoney
		}
	}
	info := miner.Info()
	if c.AllMiners[info.MinerType] == nil {
		c.AllMiners[info.MinerType] = make(map[uuid.UUID]miners.Miner)
	}
	c.AllMiners[info.MinerType][info.ID] = miner

	coalCh := miner.Run(c.CompanyCtx)
	go func() {
		for v := range coalCh {
			c.IncomeChannel <- v
		}
		c.mtx.Lock()
		delete(c.AllMiners[info.MinerType], info.ID)
		c.mtx.Unlock()
	}()
	c.Statistic.AllMiners[info.MinerType]++
	return miner, nil
}

func NewCompany(ctx context.Context) *Company {
	context, cancel := context.WithCancel(ctx)
	c := &Company{
		AllMiners:     make(map[string]map[uuid.UUID]miners.Miner),
		IncomeChannel: make(chan int64),
		Statistic:     NewCompanyStatistics(),
		CompanyCtx:    context,
		CompanyStop:   cancel,
		Equipment:     *NewEquipment(),
		TimeStart:     time.Now(),
		TimeComplete:  nil,
		GameTime:      nil,
	}
	go c.permanentIncome()
	go c.collectIncome()

	return c
}

func (c *Company) collectIncome() {
	for {
		select {
		case <-c.CompanyCtx.Done():
			return
		case income := <-c.IncomeChannel:
			c.Statistic.Balance.Add(income)
			c.Statistic.TotalEarned.Add(income)
		}
	}
}
func (c *Company) permanentIncome() {
	for {
		select {
		case <-c.CompanyCtx.Done():
			return
		case <-time.After(1 * time.Second):
			c.IncomeChannel <- 1
		}
	}
}
