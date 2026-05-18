package http

import (
	"context"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

type HTTPserver struct {
	httpHandlers *HTTPhandlers
	Server       *http.Server
	httpContext  context.Context
}

func NewHTTPserver(handlers *HTTPhandlers, ctx context.Context) *HTTPserver {
	return &HTTPserver{
		httpHandlers: handlers,
		httpContext:  ctx,
	}
}
func (h *HTTPserver) StartServer() error {
	router := mux.NewRouter()
	router.Path("/miners").Methods(http.MethodPost).HandlerFunc(h.httpHandlers.CreateMinerHandler)
	router.Path("/miners").Methods(http.MethodGet).HandlerFunc(h.httpHandlers.GetMinersHandler)
	router.Path("/miners/salaries").Methods(http.MethodGet).HandlerFunc(h.httpHandlers.GetMinersSalariesHandler)

	router.Path("/equipment").Methods(http.MethodPost).HandlerFunc(h.httpHandlers.BuyNewEquipment)
	router.Path("/equipment/prices").Methods(http.MethodGet).HandlerFunc(h.httpHandlers.GetEquipPrices)
	router.Path("/equipment").Methods(http.MethodGet).HandlerFunc(h.httpHandlers.CheckEquipmentHandler)

	router.Path("/company").Methods(http.MethodGet).HandlerFunc(h.httpHandlers.CheckStatsHandler)
	router.Path("/company/complete").Methods(http.MethodPost).HandlerFunc(h.httpHandlers.CompleteGameHandlers)

	server := http.Server{
		Addr:    ":9091",
		Handler: router,
	}
	go func() {
		<-h.httpContext.Done()
		server.Shutdown(context.Background())
	}()
	err := server.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		return nil
	}

	return err
}
