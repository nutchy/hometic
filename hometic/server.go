package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Pair struct {
	Device uint64
	UserID uint64
}

// git init, git add server.go go.mod, git commit -m "[Nong] init project"
func main() {
	fmt.Println("hello hometic : I'm Gopher!!")
	r := mux.NewRouter()
	r.HandleFunc("/pair-device", PairDeviceHandler).Methods(http.MethodPost)
	addr := fmt.Sprintf("0.0.0.0:%s", os.Getenv("PORT"))
	fmt.Println("Addr:", addr)
	server := http.Server{
		Addr:    addr,
		Handler: r,
	}
	log.Println("starting...")
	log.Fatal(server.ListenAndServe())
}

func PairDeviceHandler(w http.ResponseWriter, r *http.Request) {
	var p Pair
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	defer r.Body.Close()
	fmt.Printf("pair %+v", p)
	w.Write([]byte(`{"status":"active"}`))
}
