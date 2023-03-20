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

	decoder := json.NewDecoder(r.Body)
	var reqBody map[string]interface{}
	decoder.Decode(&reqBody)
	action := reqBody["action"].(string)
	test, _ := json.Marshal(map[string]string{"responseText": commandHandler(action)})
	w.Write(test)
}
