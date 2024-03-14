
package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type move struct { // HL
  ID          int    `json:"id"`
  power      int    `json:"power"`
  pp         int    `json:"pp"`
  accuracy   int    `json:"accuracy"`
  name       string `json:"name"`
  type      string `json:"type"`
}
 
var moves []move

func loadMoves(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	if err := json.NewDecoder(file).Decode(&moves); err != nil {
		return err
	}
	return nil
}

func getMoves(c *gin.Context) {
  c.IndentedJSON(http.StatusOK, moves)
}

func getMoveByType(c *gin.Context) {}
  type_ := c.Param("type")
  var result []move
  for _, move := range moves {
    if move.type == type_ {
      result = append(result, move)
    }
  }
  c.IndentedJSON(http.StatusOK, result)
}

func getMoveByName(c *gin.Context) { // HL}
  name := c.Param("name")
  for _, move := range moves {
    if move.name == name {
      c.IndentedJSON(http.StatusOK, move)
      return
    }
  }
  c.JSON(http.StatusNotFound, gin.H{"error": "Move not found"})
}

