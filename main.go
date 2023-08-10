package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var fortunes = []string{
	"Dai-kichi",
	"Kichi",
	"Chuu-kichi",
	"Sho-kichi",
	"Sue-kichi",
	"Kyo",
	"Dai-kyo",
}

type Response struct {
	Fortune   string `json:"fortune"`
	Health    string `json:"health"`
	Residence string `json:"residence"`
	Travel    string `json:"travel"`
	Study     string `json:"study"`
	Love      string `json:"love"`
}

var currentTime = func() time.Time {
	return time.Now()
}

func handler(w http.ResponseWriter, r *http.Request) {
	var fortune string

	now := currentTime()
	if now.Month() == 1 && (1 <= now.Day() && now.Day() <= 3) {
		fortune = "Dai-kichi"
	} else {
		fortune = fortunes[rand.Intn(len(fortunes))]
	}

	response := Response{
		Fortune:   fortune,
		Health:    "You will fully recover, but stay attentive after you do.",
		Residence: "You will have good fortune with a new house.",
		Travel:    "When traveling, you may find something to treasure.",
		Study:     "Things will be better. It may be worth aiming for a school in a different area.",
		Love:      "The person you are looking for is very close to you.",
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		fmt.Fprintln(os.Stderr, "Usage: omikuji <port>")
		os.Exit(1)
	}
	http.HandleFunc("/", handler)

	err := http.ListenAndServe(":"+args[0], nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
