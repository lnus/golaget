package golaget

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestSiteSearch(t *testing.T) {
	godotenv.Load()
	key := os.Getenv("PRIMARY_KEY")
	i, err := New(key, false)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	r, err := i.SiteSearch("UPPSALA")
	if len(*r) <= 0 || err != nil {
		t.Log(len(*r), err)
		t.Fail()
	}
}
