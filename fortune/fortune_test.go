package fortune_test

import (
	"github.com/shima8823/Omikuji-API/fortune"
	"testing"
	"time"
)

func TestGetFortune(t *testing.T) {
	t.Parallel()
	fortune.CurrentTime = func() time.Time {
		return time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	}

	if fortune := fortune.GetFortune(); fortune != "Dai-kichi" {
		t.Errorf("Expected 'Dai-kichi' but got %s", fortune)
	}

	fortune.CurrentTime = func() time.Time {
		return time.Date(2023, 1, 4, 0, 0, 0, 0, time.UTC)
	}

	isFound := false
	fortuneTitle := fortune.GetFortune()
	for _, v := range fortune.Fortunes {
		if fortuneTitle == v {
			isFound = true
			break
		}
	}
	if !isFound {
		t.Errorf("Expected one of %v but got %s", fortune.Fortunes, fortuneTitle)
	}
}

func TestGetResponse(t *testing.T) {
	t.Parallel()
	resp := fortune.GetResponse()
	if len(resp.Fortune) == 0 {
		t.Errorf("empty fortune")
	}
	if len(resp.Health) == 0 {
		t.Errorf("empty health")
	}
	if len(resp.Love) == 0 {
		t.Errorf("empty love")
	}
	if len(resp.Residence) == 0 {
		t.Errorf("empty residence")
	}
	if len(resp.Study) == 0 {
		t.Errorf("empty study")
	}
	if len(resp.Travel) == 0 {
		t.Errorf("empty travel")
	}
}
