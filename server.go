package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"fmt"
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
	result, err := json.Marshal(map[string]string{"responseText": commandHandler(reqBody["action"].(string))})
	if err != nil {
		fmt.Println("Error marshalling response JSON in requestHandler(): ", err)
		return
	}
	w.Write(result)
}

func helpHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Add("Access-Control-Allow-Methods", "POST")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json")

	content, err := ioutil.ReadFile("json/help.json")
	if(err != nil) {
		fmt.Println("Error reading help.json in helpHandler(): ", err)
		return
	}

	w.Write(content)
}
