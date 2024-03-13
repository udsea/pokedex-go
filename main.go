package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := loadPokemons("pokedex.json"); err != nil {
		log.Fatalf("Failed to load pokemons: %v", err)
	}
	r := gin.Default()
	r.GET("/pokemons", getPokemons)
	r.GET("/pokemons/:id", getPokemonByID)
	r.Run("localhost:8080")
}
