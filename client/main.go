package main

import (
	"gitbub.com/eminoz/graceful-fiber/client/config"
	"gitbub.com/eminoz/graceful-fiber/client/router"
)

func main() {
	config.SetupConfig()
	r := router.UserRoute()
	r.Listen(":4041")

}
