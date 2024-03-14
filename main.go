package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := loadPokemons("pokedex.json"); err != nil {
		log.Fatalf("Failed to load pokemons: %v", err)
	}

	if err := loadMoves("moves.json"); err != nil {
		log.Fatalf("Failed to load moves: %v", err)
	}

	r := gin.Default()
	r.GET("/api/v1/pokemons", getPokemons)
	r.GET("/api/v1/pokemons/:id", getPokemonByID)
	r.GET("/api/v1/moves", getMoves)
	r.GET("/api/v1/moves/:type", getMovesByType)
	r.Run("localhost:8080")
}
