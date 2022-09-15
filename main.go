package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world!"))
	}).Methods(http.MethodGet)

	port := os.Getenv("PORT")
	log.Println("Listening on port", port)
	if IsProduction() {
		log.Println("🏭 Running in production mode...")
		log.Fatal(http.ListenAndServeTLS(":"+port, "./cert.pem", "./privkey.pem", r))
	} else {
		log.Fatal(http.ListenAndServe(":"+port, r))
	}
}

func IsProduction() bool {
	return os.Getenv("ENV") == "production"
}
