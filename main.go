package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"1inch-test/api"
	"1inch-test/contarcts/multicall"
	"1inch-test/contarcts/univ2"
	"1inch-test/core"
	"1inch-test/pool"

	"github.com/gorilla/mux"
)

var (
	address = "127.0.0.1:8080"
)

func main() {
	prividerURL, ok := os.LookupEnv("PROVIDER_URL")
	if !ok {
		log.Fatal("INFURA_URL not found")
	}

	poolContract, err := univ2.NewContract()
	if err != nil {
		log.Fatal("init univ2 contract failed", err)
	}

	multicallContarct, err := multicall.NewContract(prividerURL)
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