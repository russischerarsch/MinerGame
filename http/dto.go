package http

type MinersSalariesDTO struct {
	LittleSalary int
	MiddleSalary int
	LeadSalary   int
}

func NewMinersSalaries(little, middle, lead int) *MinersSalariesDTO {
	return &MinersSalariesDTO{
		LittleSalary: little,
		MiddleSalary: middle,
		LeadSalary:   lead,
	}
}

type EquipmentPrices struct {
	Pickaxe  int
	Vents    int
	Trolleys int
}

func NewEquipPrices(axes, vents, trolleys int) *EquipmentPrices {
	return &EquipmentPrices{
		Pickaxe:  axes,
		Vents:    vents,
		Trolleys: trolleys,
	}
}
