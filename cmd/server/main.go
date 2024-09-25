package main

import (
	"api-go-full/configs"
	_ "api-go-full/docs"
	"api-go-full/internal/entity"
	"api-go-full/internal/infra/db"
	"api-go-full/internal/infra/webserver/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
)

//@title API Go Full
//@version 1.0
//@description API para gerenciamento de produtos
//@termsOfService http://swagger.io/terms/

//@contact.name Marco Fameli
//@contact.url http://www.swagger.io/support
//@contact.email email@example.com

//@license.name Teste
//@license.url http://www.swagger.io/support

//@host localhost:8080
//@BasePath /

//@securityDefinitions.apikey ApiKeyAuth
//@in header
//@name Authorization

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		log.Fatal("Erro ao carregar configuração", err)
	}
	// arrumei o dbConn para funcionar
	dbConn, err := gorm.Open(postgres.Open("host=db user=postgres password=admin dbname=teste port=5432 sslmode=disable"), &gorm.Config{})
	if err != nil {
		panic("ERRO AO CONECTAR NO BANCO DE DADOS" + err.Error())
	}

	err = dbConn.AutoMigrate(&entity.User{}, &entity.Product{}) // Certifique-se de ter as entidades corretas
	if err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}

	productDB := db.NewProduct(dbConn)
	productHandler := handlers.NewProductHandler(productDB)

	UserDB := db.NewUser(dbConn)
	userHandler := handlers.NewUseHandler(UserDB)

	r := chi.NewRouter()
	r.Use(LogRequest)
	r.Use(middleware.WithValue("jwt", configs.TokenAuth))
	r.Use(middleware.WithValue("JwtExpiresIn", configs.JwtExpiresIn))

	r.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(configs.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Post("/", productHandler.CreateProduct)
		r.Get("/", productHandler.GetProducts)
		r.Get("/{id}", productHandler.GetProduct)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
	})
	r.Post("/users", userHandler.Create)
	r.Post("/users/generate_jwt", userHandler.GetJWT)

	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8080/docs/doc.json")))

	http.ListenAndServe(":8080", r)

}
func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Context().Value("user")
		log.Println(r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
