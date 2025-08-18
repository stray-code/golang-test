package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type EvolutionChain struct {
	Chain Chain `json:chain`
}

type Chain struct {
	Species Species `json:species`
}

type Species struct {
	Name string `json:name`
	Url  string `json:url`
}

func main() {
	resp, err := http.Get("https://pokeapi.co/api/v2/evolution-chain/10")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var data EvolutionChain

	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println(data.Chain.Species.name)
}
