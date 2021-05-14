package golaget

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestAgentGet(t *testing.T) {
	godotenv.Load()
	key := os.Getenv("PRIMARY_KEY")
	i, err := New(key, true)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	r, err := i.AgentGet()
	if len(*r) <= 0 || err != nil {
		t.Log(len(*r), err)
		t.Fail()
	}
}

func TestAgentGetById(t *testing.T) {
	godotenv.Load()
	key := os.Getenv("PRIMARY_KEY")
	i, err := New(key, true)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	r, err := i.AgentGetById("012003-91A")
	if r.City != "SANDHAMN" || err != nil {
		t.Log(r.City, err)
		t.Fail()
	}
}
