package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/kevinmalo/Costanera700/internal/models"
	"github.com/kevinmalo/Costanera700/migrations/upload_buyers"
	"github.com/kevinmalo/Costanera700/migrations/upload_products"
	"github.com/kevinmalo/Costanera700/migrations/upload_transactions"
	"net/http"
	"time"
)

const (
	layoutISO = "2006-01-02"
)

func Routes() *chi.Mux {
	mux := chi.NewMux()

	//globals middlewares
	mux.Use(
		middleware.Logger,    // log every http request
		middleware.Recoverer, // recover if a panic occurs
	)

	//Routes
	//Add all data in db
	mux.Post("/uploads/{date}", uploadsHandler)
	//Return product history by id
	mux.Get("/transaction/{id}", shoppingHistoryHandler)
	//Get all buyers
	mux.Get("/buyers", buyerHandler)
	//Get buyers with same ip
	mux.Get("/buyers-ip/{id}", buyerIpHandler)
	//Get best sellers
	mux.Get("/best-sellers", bestSellerHandler)

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

func buyerIpHandler(w http.ResponseWriter, r *http.Request) {

	buyerId := chi.URLParam(r, "id")

	// Search buyers ips
	buyersIds := models.GetTransactionsIp(buyerId)

	//Search buyers names
	buyersNames := models.GetBuyerName(buyersIds)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(buyersNames)

}

func uploadsHandler(w http.ResponseWriter, r *http.Request) {

	date := chi.URLParam(r, "date")
	t, _ := time.Parse(layoutISO, date)
	timeUnix := int(t.Unix())


	//// Set all
	upload_buyers.SetBuyers(timeUnix)
	upload_products.SetBuyers(timeUnix)
	upload_transactions.SetTransactions(timeUnix)

}

func bestSellerHandler(w http.ResponseWriter, r *http.Request) {

	bestSellers := models.GetBestSellers()

	//// Set all
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(bestSellers)
}
