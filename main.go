package main

import (
	"crudapp/routes"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	fmt.Println("Starting server.")
	godotenv.Load(".env")

	// * Router.
	router := mux.NewRouter()
	fs := http.FileServer(http.Dir("./public/"))
	router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", fs))
	routes.RegisterPersonRoutes(router)

	http.Handle("/", router)
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowCredentials: true,
		AllowedMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
	})
	handler := c.Handler(router)
	port := os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(":"+port, handler))
}