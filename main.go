package main

import (
	"project/config"
	"project/src"
)

var conf = config.New()

func main() {
	server := src.App()
	server.Run(conf.ADDRESS)
}
