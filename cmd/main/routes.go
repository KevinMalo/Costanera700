package main

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/kevinmalo/Costanera700/internal/models"
	"net/http"
)

func Routes() *chi.Mux {
	mux := chi.NewMux()

	//globals middlewares
	mux.Use(
		middleware.Logger, // log every http request
		middleware.Recoverer, // recover if a panic occurs
		)

	//BuyerRoute
	mux.Post("/buyers", nil)
	mux.Get("/buyers", buyerHandler)

	//TEST
	mux.Get("/hello", helloHandler)

	return mux

}

func helloHandler(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("done-by","kevin")

	res := map[string]interface{}{"message":"hello world"}

	_ = json.NewEncoder(w).Encode(res)
}

func buyerHandler(w http.ResponseWriter, r *http.Request)  {

	buyers := models.GetBuyers()

	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Access-Control-Allow-Origin","*")
	w.Write(buyers)

}