package main

import (
	"project-skbackend/configs"
	"project-skbackend/internal/apps"
)

func main() {
	// * get and map the env value to an object
	cfg := configs.GetInstance()
	apps.Run(cfg)
}
