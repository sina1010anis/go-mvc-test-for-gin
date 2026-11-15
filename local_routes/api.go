package local_routes

import (
	"web/controllers"

	"github.com/gin-gonic/gin"
)

var r = gin.Default()

func webRoute() {

	r.GET("/t1/:num1/:num2/:ope", controllers.T1)
	r.GET("/users", controllers.T2)

	// {

	// 	r2 := r.Group("t2")
	// 	r2.GET("/")

	// }

} // is all youre route

func runRoute() {

	r.Run()

}

func BuildRoute() {

	webRoute()

	runRoute()

}
