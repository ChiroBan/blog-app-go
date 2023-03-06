package main

import (

	controller "github.com/pilinux/gorest/controller"
	model "github.com/pilinux/gorest/database"
)

func main() {
	model.Init()
	controller.Start()
}