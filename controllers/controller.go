package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func T1(c *gin.Context) {

	num1, err1 := strconv.Atoi(c.Param("num1"))
	num2, err2 := strconv.Atoi(c.Param("num2"))
	var res int

	if err1 != nil || err2 != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"Status": "Error",
			"Body": gin.H{
				"URL":       c.Request.URL.Path,
				"Message":   "Pleas type only inter",
				"Exeption1": err1,
				"Exeption2": err2,
				"Data":      nil,
			},
		})

		return

	}

	ope := c.Param("ope")

	switch ope {

	case "+":

		res = num1 + num2
	case "-":
		res = num1 - num2
	case "*":

		res = num1 * num2
	case "dir":
		res = num1 / num2
	default:
		c.JSON(http.StatusBadRequest, gin.H{
			"Status": "Error",
			"Body": gin.H{
				"URL":     c.Request.URL.Path,
				"Message": "Pleas type standard opertot (+ , - , * , dir)",
				"Data":    nil,
			},
		})

		return
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"Status": "OK",
		"Body": gin.H{
			"URL":  c.Request.URL.Path,
			"Data": res,
		},
	})

}
