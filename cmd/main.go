package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/lnus/systemgolaget"
)

func perror(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	godotenv.Load()
	api := os.Getenv("PRIMARY_KEY")
	fmt.Println(api)
	c, err := systemgolaget.New(api, true)
	perror(err)

	foo, err := c.AgentGet()
	perror(err)

	for k, v := range *foo {
		fmt.Println(k)
		fmt.Println(v.Name)
		fmt.Println(v.IsOpen)
	}

	bar, err := c.AgentGetById("")
	perror(err)
	fmt.Println(bar.Name)
	fmt.Println(bar.Address)
	fmt.Println(bar.AgentID)
	fmt.Println(bar.IsOpen)
	fmt.Println(bar.AgentPickupHours)
}
