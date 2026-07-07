package main

import (
	"github.com/rafaelsouzaribeiro/climate-by-zip-code-exercise-in-golang/configs"
	"github.com/rafaelsouzaribeiro/climate-by-zip-code-exercise-in-golang/internal/infra/di"
)

func main() {
	_, _ = configs.LoadConfig(".")
	server := di.NewDi()

	if err := server.Start(); err != nil {
		panic(err)
	}
}
