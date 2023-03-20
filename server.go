package main

import (
	"os"
	"net/http"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	serverPort := os.Getenv("PORT")

	http.HandleFunc("/", requestHandler)
	http.ListenAndServe(":"+serverPort, nil)
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Add("Access-Control-Allow-Methods", "POST")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte(`{"responseText": "Hello World"}`))
}