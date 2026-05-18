package company

import "fmt"

const (
	Pickaxe       = "pickaxe"
	Vents         = "vents"
	Trolleys      = "trolleys"
	PickaxesPrice = 3000
	VentsPrice    = 15_000
	TrolleysPrice = 50_000
)

type Equipment struct {
	Pickaxes bool
	Vents    bool
	Trolleys bool
}

func NewEquipment() *Equipment {
	return &Equipment{}
}
func (e *Equipment) BuyPickaxes() {
	e.Pickaxes = true
	fmt.Println("Pickaxes are purchased")
}
func (e *Equipment) BuyVents() {
	e.Vents = true
	fmt.Println("Ventilaitions are purchased")
}
func (e *Equipment) BuyTrolleys() {
	e.Trolleys = true
	fmt.Println("Trolleys are purchased")
}
func (e Equipment) IsPickaxes() bool {
	return e.Pickaxes
}
func (e Equipment) IsVents() bool {
	return e.Vents
}
func (e Equipment) IsTrolleys() bool {
	return e.Trolleys
}
func (e Equipment) AllBought() bool {
	return e.IsPickaxes() && e.IsVents() && e.IsTrolleys()
}
