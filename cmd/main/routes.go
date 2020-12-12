package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/kevinmalo/Costanera700/internal/models"
	"github.com/kevinmalo/Costanera700/migrations/upload_buyers"
	"github.com/kevinmalo/Costanera700/migrations/upload_products"
	"github.com/kevinmalo/Costanera700/migrations/upload_transactions"
	"net/http"
	"strconv"
)

func Routes() *chi.Mux {
	mux := chi.NewMux()

	//globals middlewares
	mux.Use(
		middleware.Logger,    // log every http request
		middleware.Recoverer, // recover if a panic occurs
	)

	//Routes
	mux.Post("/uploads/{id}", uploadsHandler)
	mux.Get("/buyers", buyerHandler)

	return mux

}

func buyerHandler(w http.ResponseWriter, r *http.Request) {

	buyers := models.GetBuyers()

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(buyers)

}

func uploadsHandler(w http.ResponseWriter, r *http.Request) {

	date := chi.URLParam(r, "id")
	i, _ := strconv.Atoi(date)

	//// Set all
	upload_buyers.SetBuyers(i)
	upload_products.SetBuyers(i)
	upload_transactions.SetTransactions(i)

}
