package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PokemonInfo struct {
	PokemonStats
	PokemonForms
}
type PokemonStats struct {
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
}
type PokemonForms struct {
	Name    string `json:"name"`
	Sprites struct {
		BackDefault      string `json:"back_default"`
		BackFemale       string `json:"back_female"`
		BackShiny        string `json:"back_shiny"`
		BackShinyFemale  string `json:"back_shiny_female"`
		FrontDefault     string `json:"front_default"`
		FrontFemale      string `json:"front_female"`
		FrontShiny       string `json:"front_shiny"`
		FrontShinyFemale string `json:"front_shiny_female"`
	} `json:"sprites"`
}

type Data struct {
	ID int `json:"id"`
}

func mergePokemonInfo(pkForm PokemonForms, pk PokemonStats) PokemonInfo {
	return PokemonInfo{
		PokemonStats: pk,
		PokemonForms: pkForm,
	}
}

func createPokemonRequest(c *gin.Context) {
	var data Data
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var id = strconv.Itoa(data.ID)
	pk, err := callPokemonAPI(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	pkForm, err := callPokemonFormAPI(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	pkInfo := mergePokemonInfo(pkForm, pk)
	c.JSON(http.StatusOK, pkInfo)
}

func callPokemonFormAPI(id string) (PokemonForms, error) {

	var pokemonForm PokemonForms
	url := "https://pokeapi.co/api/v2/pokemon-form/" + id

	resp, err := http.Get(url)
	if err != nil {
		return pokemonForm, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&pokemonForm); err != nil {
		return pokemonForm, err
	}

	return pokemonForm, nil

}
func callPokemonAPI(id string) (PokemonStats, error) {

	var pokemon PokemonStats
	url := "https://pokeapi.co/api/v2/pokemon/" + id

	resp, err := http.Get(url)
	if err != nil {
		return pokemon, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&pokemon); err != nil {
		return pokemon, err
	}

	return pokemon, nil

}
func main() {

	router := gin.Default()
	router.POST("/pokemon", createPokemonRequest)
	router.Run("localhost:8080")
}
