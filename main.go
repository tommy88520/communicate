package main

import (
	"ginchat/router"
	"ginchat/utils"
)

func main() {
	r := router.Router()
	utils.InitConfig()
	utils.InitMySQL()

	r.Run(":80") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
