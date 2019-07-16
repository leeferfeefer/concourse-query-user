package main

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var ctx, cancel = context.WithCancel(context.Background())

type DTO struct {
	Data Data `json:"data"`
}
type Data struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func handleSubstance(w http.ResponseWriter, r *http.Request) {

	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "DELETE, POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	saveData(w, r)
}

func saveData(w http.ResponseWriter,r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "cannot read body", http.StatusBadRequest)
		panic(err)
	}

	err = ioutil.WriteFile("output.txt", body, os.ModePerm)
	if err != nil {
		log.Printf("Error creating output.txt: %v", err)
		http.Error(w, "Cannot read body", http.StatusBadRequest)
		panic(err)
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE, POST, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
	w.WriteHeader(http.StatusOK)

	cancel()
}


func main() {
	http.HandleFunc("/substance", handleSubstance)

	server := &http.Server{Addr: ":3001"}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Printf("Httpserver: ListenAndServe() error: %s", err)
		}
	}()
	<-ctx.Done()
	if err := server.Shutdown(ctx); err != nil && err != context.Canceled {
		log.Println(err)
	}
}