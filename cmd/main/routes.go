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

//Chi routes
func Routes() *chi.Mux {
	mux := chi.NewMux()

	//Globals middlewares
	mux.Use(
		middleware.Logger,    // log every http request
		middleware.Recoverer, // recover if a panic occurs
	)

	//URL paths and handlers
	//Get all buyers
	mux.Get("/buyers", buyerHandler)
	//Add all data in db
	mux.Post("/uploads/{date}", uploadsHandler)
	//Get product history by id
	mux.Get("/transaction/{id}", purchaseHistoryHandler)
	//Get buyers with same ip
	mux.Get("/buyers-ip/{id}", buyerIpHandler)
	//Get best sellers
	mux.Get("/best-sellers", bestSellerHandler)

	return mux

}

//Get all buyers in db
func buyerHandler(w http.ResponseWriter, r *http.Request) {

	buyers := models.GetBuyers()

	// Set headers
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(buyers)

}

//Get buyers purchase history
func purchaseHistoryHandler(w http.ResponseWriter, r *http.Request) {

	buyerId := chi.URLParam(r, "id")

	//Search products id data
	productsIds := models.GetTransactionsHistory(buyerId)

	//Search products names
	productsNames := models.GetProductsNames(productsIds)

	// Set headers
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(productsNames)

}

//Get buyers with same ip
func buyerIpHandler(w http.ResponseWriter, r *http.Request) {

	buyerId := chi.URLParam(r, "id")

	// Search buyers ips
	buyersIds := models.GetTransactionsIp(buyerId)

	//Search buyers names
	buyersNames := models.GetBuyerName(buyersIds)

	// Set headers
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(buyersNames)

}

//Get buyers with same ip
func uploadsHandler(w http.ResponseWriter, r *http.Request) {

	//Get and parse date to Unix Timestamp
	date := chi.URLParam(r, "date")
	t, _ := time.Parse(layoutISO, date)
	timeUnix := int(t.Unix())

	//Process all data and put in db
	upload_buyers.SetBuyers(timeUnix)
	upload_products.SetProducts(timeUnix)
	upload_transactions.SetTransactions(timeUnix)

}

//Get the 5 best-selling products
func bestSellerHandler(w http.ResponseWriter, r *http.Request) {

	bestSellers := models.GetBestSellers()

	// Set headers
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(bestSellers)
}
