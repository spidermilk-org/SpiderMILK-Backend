package main

import (
	"os"
	"net/http"
	"encoding/json"
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

	var reqBody map[string]interface{}
	json.NewDecoder(r.Body).Decode(&reqBody)
	result, _ := json.Marshal(map[string]string{"responseText": commandHandler(reqBody["action"].(string))})
	w.Write(result)
}
