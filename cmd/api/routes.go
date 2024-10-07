package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

/* returns http.Handler*/
func (app *Config) routes() http.Handler {
	mux := chi.NewRouter()

	//specify who is allowed to connect
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	mux.Get("/api/v1/getusers", app.GetUsers)
	// Add the Prometheus metrics endpoint to the router
	mux.Handle("/metrics", promhttp.Handler())

	//POS Mainnet & Testnet: Missed Checkpoint
	mux.Get("/api/v1/pos/mainnet/mainnet-missed-checkpoint", app.MainnetMissedCheckpoint)
	mux.Get("/api/v1/pos/testnet/mainnet-missed-checkpoint", app.TestnetMissedCheckpoint)

	//POS Mainnet & Testnet: Heimdell Block Height
	mux.Get("/api/v1/pos/mainnet/heimdal-block-height", app.MainnetHeimdalBlockHeight)
	mux.Get("/api/v1/pos/testnet/heimdal-block-height", app.TestnetHeimdalBlockHeight)

	//POS Mainnet &Testnet: Bor Latest Block Detail
	mux.Get("/api/v1/pos/mainnet/bor-latest-block-details", app.MainnetBorLatestBlockDetails)
	mux.Get("/api/v1/pos/testnet/bor-latest-block-details", app.TestnetBorLatestBlockDetails)

	//POS Mainnet &Testnet: State Sync
	mux.Get("/api/v1/pos/mainnet/state-sync", app.MainnetStateSync)
	mux.Get("/api/v1/pos/testnet/state-sync", app.TestnetStateSync)

	return mux
}
