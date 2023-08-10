package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/shima8823/Omikuji-API/fortune"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	response := fortune.GetResponse()

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
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
	http.HandleFunc("/", handler)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
