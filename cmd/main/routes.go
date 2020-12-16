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
	mux.Get("/transaction/{id}", shoppingHistoryHandler)
	mux.Get("/buyers", buyerHandler)
	mux.Get("/buyers/{id}", buyerIdHandler)

	return mux

}

func buyerHandler(w http.ResponseWriter, r *http.Request) {

	buyers := models.GetBuyers()

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(buyers)

}

func shoppingHistoryHandler(w http.ResponseWriter, r *http.Request) {

	buyerId := chi.URLParam(r, "id")

	//Search products id data
	productsIds := models.GetTransactionsHistory(buyerId)

	//Search products names
	productsNames := models.GetProductsNames(productsIds)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(productsNames)

}

func buyerIdHandler(w http.ResponseWriter, r *http.Request) {

	buyerId := chi.URLParam(r, "id")

	// Search buyer data
	buyer := models.GetBuyersById(buyerId)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(buyer)

}

func uploadsHandler(w http.ResponseWriter, r *http.Request) {

	date := chi.URLParam(r, "id")
	i, _ := strconv.Atoi(date)

	//// Set all
	upload_buyers.SetBuyers(i)
	upload_products.SetBuyers(i)
	upload_transactions.SetTransactions(i)

}
