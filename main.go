package main

import (
	"log"
	"os"
	"server-info/app"
)

func main() {
	application := app.Generate()
	if erro := application.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}
