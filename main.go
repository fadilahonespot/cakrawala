package main

import (
	"github.com/fadilahonespot/cakrawala/server/container"
	"github.com/joho/godotenv"
)

func main() {
	// Load Env
	godotenv.Load()

	// Setup Container
	container.NewContainer()
}