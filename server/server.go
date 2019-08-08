package main

import (
	"context"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var ctx, kill = context.WithCancel(context.Background())

type DTO struct {
	Data struct {
		Username string `json:"username"`
		Password string `json:"password"`
	} `json:"data"`
}

func handleToll(w http.ResponseWriter, r *http.Request) {

	if r.Method == "OPTIONS" {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		}
		w.Header().Set("Access-Control-Allow-Methods", "DELETE, POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err := saveData(w, r)
	if err != nil {
		handleBadRequest(w, err, "Could not save data")
		return
	}
	w.WriteHeader(http.StatusOK)

	kill()
}

func saveData(w http.ResponseWriter,r *http.Request) error {

	body, err := ioutil.ReadAll(r.Body)
	handleErr(err, "Could not read body")

	var dto DTO
	err = json.Unmarshal(body, &dto)
	handleErr(err, "Could not unmarshal body")

	var decodedUsername, decodedPassword []byte
	decodedUsername, err = b64.StdEncoding.DecodeString(dto.Data.Username)
	handleErr(err, "Could not decode username")

	decodedPassword, err = b64.StdEncoding.DecodeString(dto.Data.Password)
	handleErr(err, "Could not decode password")

	dto.Data.Username = string(decodedUsername)
	dto.Data.Password = string(decodedPassword)

	body, err = json.Marshal(dto)
	handleErr(err, "Could not marshal data")

	err = ioutil.WriteFile("output.txt", body, os.ModePerm)
	handleErr(err, "Could not write data to file")

	return err
}

func handleBadRequest(w http.ResponseWriter, err error, message string) {
	http.Error(w, fmt.Sprintf("Error: %s - %s", err, message), http.StatusBadRequest)
}

func handleErr(err error, message string) {
	if err != nil {
		log.Printf("Error: %s - %s", err, message)
	}
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("../app/dist/")))
	http.HandleFunc("/toll", handleToll)

	server := &http.Server{Addr: ":3001"}

	go func() {
		if err := server.ListenAndServeTLS("server.crt", "server.key"); err != nil && err != http.ErrServerClosed  {
			log.Printf("Httpserver: ListenAndServe() error: %s", err)
		}
	}()
	<-ctx.Done()
	if err := server.Shutdown(ctx); err != nil && err != context.Canceled {
		log.Println(err)
	}
}
