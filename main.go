package main

import (
	"context"
	"fmt"
	"secPetProject/company"
	"secPetProject/http"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	company := company.NewCompany(ctx)
	handlers := http.NewHTTPHandlers(company)
	server := http.NewHTTPserver(handlers, ctx)

	fmt.Println("game started!")
	if err := server.StartServer(); err != nil {
		fmt.Println("error while running HTTP server:", err)
	}
}
