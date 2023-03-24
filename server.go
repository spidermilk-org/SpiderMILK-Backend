package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	serverPort := os.Getenv("PORT")

	http.HandleFunc("/", requestHandler)
	http.HandleFunc("/help", helpHandler)
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

func helpHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Add("Access-Control-Allow-Methods", "POST")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json")

	content, _ := ioutil.ReadFile("json/help.json")

	w.Write(content)
}
