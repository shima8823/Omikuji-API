package fortune

import (
	"math/rand"
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

var CurrentTime = func() time.Time {
	return time.Now()
}

func GetFortune() string {
	now := CurrentTime()
	if now.Month() == 1 && (1 <= now.Day() && now.Day() <= 3) {
		return "Dai-kichi"
	}
	return fortunes[rand.Intn(len(fortunes))]
}

func GetResponse() Response {
	return Response{
		Fortune:   GetFortune(),
		Health:    "You will fully recover, but stay attentive after you do.",
		Residence: "You will have good fortune with a new house.",
		Travel:    "When traveling, you may find something to treasure.",
		Study:     "Things will be better. It may be worth aiming for a school in a different area.",
		Love:      "The person you are looking for is very close to you.",
	}
}
