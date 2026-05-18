package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"secPetProject/company"
	"secPetProject/company/miners"
)

type HTTPhandlers struct {
	company *company.Company
}

func NewHTTPHandlers(company *company.Company) *HTTPhandlers {
	return &HTTPhandlers{
		company: company,
	}
}

func (h *HTTPhandlers) CreateMinerHandler(w http.ResponseWriter, r *http.Request) {
	var MinerType string
	if err := json.NewDecoder(r.Body).Decode(&MinerType); err != nil {
		http.Error(w, NewErrorDTO(err).ErrorReadable(), http.StatusBadRequest)
		return
	}
	miner, err := h.company.HireMiner(MinerType)
	if errors.Is(err, company.UnknownMinerType) {
		http.Error(w, NewErrorDTO(err).ErrorReadable(), http.StatusBadRequest)
	} else if errors.Is(err, company.NotEnoughMoney) {
		http.Error(w, NewErrorDTO(err).ErrorReadable(), http.StatusForbidden)
	} else {
		http.Error(w, NewErrorDTO(err).ErrorReadable(), http.StatusInternalServerError)
	}
	b, err := json.MarshalIndent(miner.Info(), "", "    ")
	if err != nil {
		panic(err)
	}
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write HTTP response", err)
	}
}

func (h *HTTPhandlers) GetMinersHandler(w http.ResponseWriter, r *http.Request) {
	minerType := r.URL.Query().Get("type")
	if minerType != "" {
		MinersMap := h.company.GetMinersByType(minerType)
		b, err := json.Marshal(MinersMap)
		if err != nil {
			panic(err)
		}
		if _, err := w.Write(b); err != nil {
			fmt.Println("failed to send HTTP response", err)
		}
	}
	allMiners := h.company.GetAllMiners()
	b, err := json.Marshal(allMiners)
	if err != nil {
		panic(err)
	}
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to send HTTP response", err)
	}
}

func (h *HTTPhandlers) BuyNewEquipment(w http.ResponseWriter, r *http.Request) {
	var equipmentType string
	if err := json.NewDecoder(r.Body).Decode(&equipmentType); err != nil {
		http.Error(w, NewErrorDTO(err).ErrorReadable(), http.StatusBadRequest)
		return
	}
	equipment, err := h.company.BuyEquipment(equipmentType)
	if errors.Is(err, company.NotEnoughMoney) {
		http.Error(w, NewErrorDTO(err).ErrorReadable(), http.StatusForbidden)
	} else if errors.Is(err, company.UnknownEquipType) {
		http.Error(w, NewErrorDTO(err).ErrorReadable(), http.StatusBadRequest)
	} else {
		http.Error(w, NewErrorDTO(err).ErrorReadable(), http.StatusInternalServerError)
	}
	b, err := json.Marshal(equipment)
	if err != nil {
		panic(err)
	}
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to send http response")
		return
	}
}

func (h *HTTPhandlers) CheckEquipmentHandler(w http.ResponseWriter, r *http.Request) {
	equipment := h.company.GetEquipment()

	b, err := json.MarshalIndent(equipment, "", "    ")
	if err != nil {
		panic(err)
	}

	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write HTTP response:", err)
	}
}
func (h *HTTPhandlers) CheckStatsHandler(w http.ResponseWriter, r *http.Request) {
	stats := h.company.GetStats()

	b, err := json.MarshalIndent(stats, "", "    ")
	if err != nil {
		panic(err)
	}

	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write HTTP response:", err)
	}
}
func (h *HTTPhandlers) GetMinersSalariesHandler(w http.ResponseWriter, r *http.Request) {
	MinersSalaries := NewMinersSalaries(miners.LittleSalary, miners.MiddleSalary, miners.LeadSalary)
	b, err := json.MarshalIndent(MinersSalaries, "", "    ")
	if err != nil {
		panic(err)
	}

	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write HTTP response:", err)
	}
}
func (h *HTTPhandlers) GetEquipPrices(w http.ResponseWriter, r *http.Request) {
	equipPrices := NewEquipPrices(company.PickaxesPrice, company.VentsPrice, company.TrolleysPrice)
	b, err := json.MarshalIndent(equipPrices, "", "    ")
	if err != nil {
		panic(err)
	}

	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write HTTP response:", err)
	}
}
func (h *HTTPhandlers) CompleteGameHandlers(w http.ResponseWriter, r *http.Request) {
	if err := h.company.Complete(); err != nil {
		http.Error(w, NewErrorDTO(err).ErrorReadable(), http.StatusForbidden)
		return
	}
	b, err := json.MarshalIndent(h.company.Statistic, "", "    ")
	if err != nil {
		http.Error(w, "failed to marshal response", http.StatusInternalServerError)
		return
	}

	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write HTTP response:", err)
		return
	}
	h.company.CompanyStop()
}
