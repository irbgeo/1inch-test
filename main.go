package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"github.com/irbgeo/1inch-test/api"
	"github.com/irbgeo/1inch-test/contarcts/multicall"
	"github.com/irbgeo/1inch-test/contarcts/univ2"
	"github.com/irbgeo/1inch-test/core"
	_ "github.com/irbgeo/1inch-test/docs"
	"github.com/irbgeo/1inch-test/pool"
)

var (
	address = "127.0.0.1:8080"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading environment variables file")
	}

	providerURL := os.Getenv("PROVIDER_URL")
	if providerURL == "" {
		log.Fatal("PROVIDER_URL not found")
	}

	poolContract, err := univ2.NewContract()
	if err != nil {
		log.Fatal("init univ2 contract failed", err)
	}

	multicallContarct, err := multicall.NewContract(providerURL)
	if err != nil {
		log.Fatal("init multicall contract failed", err)
	}

	poolProvider := pool.NewProvider(poolContract, multicallContarct)

	core := core.New(poolProvider)

	r := mux.NewRouter()

	api := api.New(core)

	api.Route(r)

	srv := http.Server{
		Addr:    address,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("turn on http server failed", err)
		}
	}()

	log.Println("i'm turned on")

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch

	log.Println("goodbye")
}
