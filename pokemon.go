package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Pokemon struct {
	ID          int      `json:"id"`
	Name        Name     `json:"name"`
	Type        []string `json:"type"`
	Base        Base     `json:"base"`
	Species     string   `json:"species"`
	Description string   `json:"description"`
	Profile     Profile  `json:"profile"`
	Image       Image    `json:"image"`
}

type Name struct {
	English  string `json:"english"`
	Japanese string `json:"japanese"`
	Chinese  string `json:"chinese"`
	French   string `json:"french"`
}

type Base struct {
	HP        int `json:"HP"`
	Attack    int `json:"Attack"`
	Defense   int `json:"Defense"`
	SpAttack  int `json:"Sp. Attack"`
	SpDefense int `json:"Sp. Defense"`
	Speed     int `json:"Speed"`
}

type Profile struct {
	Height  string     `json:"height"`
	Weight  string     `json:"weight"`
	Egg     []string   `json:"egg"`
	Ability [][]string `json:"ability"`
	Gender  string     `json:"gender"`
}

type Image struct {
	Sprite    string `json:"sprite"`
	Thumbnail string `json:"thumbnail"`
	Hires     string `json:"hires"`
}

var pokemons []Pokemon

func loadPokemons(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	if err := json.NewDecoder(file).Decode(&pokemons); err != nil {
		return err
	}
	return nil
}

func getPokemons(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, pokemons)
}

func getPokemonByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	for _, pokemon := range pokemons {
		if pokemon.ID == id {
			c.JSON(http.StatusOK, pokemon)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Pokemon not found"})
}
