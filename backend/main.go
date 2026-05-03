package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type AgeRequest struct {
	Age int `json:"age"`
}

type AgeResponse struct {
	Allowed bool   `json:"allowed"`
	Message string `json:"message"`
}

func checkAgeHandler(w http.ResponseWriter, r *http.Request) {
	//cors
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req AgeRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	resp := AgeResponse{}
	if req.Age >= 18 {
		resp.Allowed = true
		resp.Message = "Доступ разрешён"
	} else {
		resp.Allowed = false
		resp.Message = "Доступ запрещён"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	http.HandleFunc("/api/check-age", checkAgeHandler)
	fmt.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
