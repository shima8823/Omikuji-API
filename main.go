package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/shima8823/Omikuji-API/fortune"
	"net/http"
	"os"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	response := fortune.GetResponse()
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func getPort() string {
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		fmt.Fprintln(os.Stderr, "Usage: omikuji <port>")
		os.Exit(1)
	}
	return args[0]
}

func main() {
	port := getPort()
	http.HandleFunc("/", Handler)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
