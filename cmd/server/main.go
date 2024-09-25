package main

import (
	"api-go-full/configs"
	"api-go-full/internal/infra/db"
	"api-go-full/internal/infra/webserver/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		log.Fatal("Erro ao carregar configuração", err)
	}
	// arrumei o dbConn para funcionar
	dbConn, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("ERRO AO CONECTAR NO BANCO DE DADOS" + err.Error())
	}
	productDB := db.NewProduct(dbConn)
	productHandler := handlers.NewProductHandler(productDB)

	UserDB := db.NewUser(dbConn)
	userHandler := handlers.NewUseHandler(UserDB, configs.TokenAuth, configs.JwtExpiresIn)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/products", productHandler.CreateProduct)
	r.Get("/products", productHandler.GetProducts)
	r.Get("/products/{id}", productHandler.GetProduct)
	r.Put("/products/{id}", productHandler.UpdateProduct)
	r.Delete("/products/{id}", productHandler.DeleteProduct)

	r.Post("/users", userHandler.Create)
	r.Post("/users/generate_jwt", userHandler.GetJWT)

	http.ListenAndServe(":8080", r)

}
