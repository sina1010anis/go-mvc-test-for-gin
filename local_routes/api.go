package local_routes

import (
	"web/controllers"

	"github.com/gin-gonic/gin"
)

var r = gin.Default()

func webRoute() {

	r.GET("/t1/:num1/:num2/:ope", controllers.T1)
	r.GET("/prime/:num", controllers.T5)

	{

		r2 := r.Group("/user")
		r2.GET("/", controllers.T2)
		r2.POST("/", controllers.AddUser)

	}

} // is all youre route

func runRoute() {

	r.Run()

}

func BuildRoute() {

	webRoute()

	runRoute()

}
