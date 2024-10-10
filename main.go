package main

import (
	"verve/controller"
	"verve/cron"
	"verve/route"
)

func main() {
	route.Init()
	controller.ExtensionInit()
	cron.Init()
}
