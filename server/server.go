package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var ctx, cancel = context.WithCancel(context.Background())

type DTO struct {
	Data string `json:"data"`
}

func handleSubstance(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	saveData(w, r)
}

func saveData(w http.ResponseWriter,r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var request DTO
	err := decoder.Decode(&request)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "cannot read body", http.StatusBadRequest)
		panic(err)
	}

	dataToWrite := []byte(request.Data)

	err = ioutil.WriteFile("output.txt", dataToWrite, os.ModePerm)
	if err != nil {
		log.Printf("Error creating output.txt: %v", err)
		http.Error(w, "Cannot read body", http.StatusBadRequest)
		panic(err)
	}

	log.Print(request.Data)

	cancel()
}


func main() {
	http.HandleFunc("/substance", handleSubstance)

	server := &http.Server{Addr: ":3000"}

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