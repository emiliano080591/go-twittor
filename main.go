package main

import (
	"log"

	"github.com/emiliano080591/go-twittor/bd"
	"github.com/emiliano080591/go-twittor/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexcion a la BD")
		return
	}

	handlers.Manejadores()
}
