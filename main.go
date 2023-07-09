package main

import (
	"lcrm2/models"
	"lcrm2/pkg/seeds"
	"lcrm2/routers"
	"log"
)

//var err error

func main() {
	models.ConnectDataBase()
	//runSeeds()
	r := routers.SetUpRouter()
	r.Run(":8080")
}

func runSeeds() {
	for _, seed := range seeds.All() {
		if err := seed.Run(models.DB); err != nil {
			log.Fatalf("Running seed '%s', failed with error: %s", seed.Name, err)
		}
	}
}
