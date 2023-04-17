package main

import (
	"ginchat/middlewares"
	routers "ginchat/router"
	"ginchat/utils"

	"github.com/gin-gonic/gin"
)

// func main() {
// 	r := router.Router()
// 	utils.InitConfig()
// 	utils.InitMySQL()
// 	utils.InitRedis()
// 	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
// }

// func testHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte(`{"message": "Hello World!"}`))
// }

// func main() {
// 	r := gin.Default()
// 	// same as
// 	// config := cors.DefaultConfig()
// 	// config.AllowAllOrigins = true
// 	// router.Use(cors.New(config))
// 	r.Use(cors.Default())

// 	r.POST("/jsonp", func(c *gin.Context) {
// 		c.String(200, "test22222")
// 	})

// 	r.Run(":8080")
// }

func main() {
	r := gin.Default()
	r.Use(middlewares.Cors())
	routers.ApiRouters(r)
	routers.SwaggerInfo(r)
	utils.InitConfig()
	utils.InitMySQL()
	utils.InitRedis()
	r.Run(":8080")

}
