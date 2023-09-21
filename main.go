package main

import (
	"hangoutin/authentication/internal/app"
)

var cfg app.Config

func init() {
	cfg.Setup()
}

func main() {
	cfg.StartServer()
}
