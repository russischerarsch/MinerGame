package company

import "errors"

var NotEnoughMoney = errors.New("Not enough money")
var UnknownMinerType = errors.New("There's no such type of miner")
var UnknownEquipType = errors.New("There's no such type of equipment")
