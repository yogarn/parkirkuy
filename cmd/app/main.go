package main

import "github.com/yogarn/parkirkuy/pkg/config"

func main() {
	config.LoadEnv()

	app := config.StartFiber()
	db := config.StartGorm()

	config.StartUp(&config.Config{
		DB:  db,
		App: app,
	})
}
