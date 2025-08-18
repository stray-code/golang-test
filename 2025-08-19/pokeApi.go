package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type EvolutionChain struct {
	Chain struct {
		Species struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"species"`
	} `json:"chain"`
}

func main() {
	resp, err := http.Get("https://pokeapi.co/api/v2/evolution-chain/10")
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("APIからエラーレスポンスが返る：%s", resp.Status)
	}

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

	fmt.Println(data.Chain.Species.Name)
}
