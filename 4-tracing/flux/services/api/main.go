package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type requestBody struct {
	Number  int               `json:"number"`
	Headers map[string]string `json:"headers"`
}

func sendReceiptHandler(w http.ResponseWriter, req *http.Request) {
	body, _ := ioutil.ReadAll(req.Body)
	var request requestBody
	if err := json.Unmarshal(body, &request); err != nil {
		log.Printf("failed to unmarshal %s\n", body)
		return
	}

	log.Printf("Number=%d TraceID=%s\n", request.Number, request.Headers["x-trace-id"])
}

func main() {
	http.HandleFunc("/", sendReceiptHandler)

	log.Println("Server running on 8081...")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		panic(err)
	}
}
