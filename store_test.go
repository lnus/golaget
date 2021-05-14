package golaget

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestStoreGet(t *testing.T) {
	godotenv.Load()
	key := os.Getenv("PRIMARY_KEY")
	i, err := New(key, false)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	r, err := i.StoreGet()
	if len(*r) <= 0 || err != nil {
		t.Log(len(*r), err)
		t.Fail()
	}
}

func TestStoreGetById(t *testing.T) {
	godotenv.Load()
	key := os.Getenv("PRIMARY_KEY")
	i, err := New(key, false)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	r, err := i.StoreGetById("0102")
	if r.City != "STOCKHOLM" || err != nil {
		t.Log(r.City, err)
		t.Fail()
	}
}
