package controllers

import (
	"net/http"
	"strconv"
	"sync"
	"web/model"
	"web/packages/prime"

	"github.com/gin-gonic/gin"
)

var nums []int
var mu sync.Mutex

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

func T2(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"data": model.UserGet(),
	})

}

func AddUser(c *gin.Context) {

	name, family, email := c.PostForm("name"), c.PostForm("family"), c.PostForm("email")

	if len(name) >= 3 && len(name) < 16 && len(family) >= 3 && len(family) < 32 && len(email) >= 7 && len(family) < 64 {

		res, err := model.UserStore(name, family, email)

		if err == nil {

			c.JSON(http.StatusCreated, gin.H{
				"data":   res,
				"Status": "Save New User",
			})

			return

		}

		c.JSON(http.StatusBadRequest, gin.H{
			"Status": "Error",
			"Body": gin.H{
				"URL":     c.Request.URL.Path,
				"Message": "For Store New User",
				"Error":   err,
				"name":    name,
				"family":  family,
				"email":   email,
			},
		})

		return

	}

	c.JSON(http.StatusBadRequest, gin.H{
		"Status": "Error",
		"Body": gin.H{
			"URL":     c.Request.URL.Path,
			"Message": "name | family | email is empty",
			"name":    name,
			"family":  family,
			"email":   email,
			"Data":    nil,
		},
	})

}
func pushJob(job chan int, num int, wg *sync.WaitGroup) {
	defer wg.Done()
	job <- num
}

func readJob(job chan int, c *gin.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range job {

		numbertestprime(val)

	}
}
func T5(c *gin.Context) {

	job := make(chan int, 10)

	var wg1 sync.WaitGroup
	var wg2 sync.WaitGroup
	nums = []int{}

	for i := 1; i <= 100; i++ {

		wg1.Add(1)

		go pushJob(job, i, &wg1)

	}

	wg2.Add(5)

	for i := 0; i < 5; i++ {

		go readJob(job, c, &wg2)
	}

	wg1.Wait()

	close(job)

	wg2.Wait()

	c.String(http.StatusOK, "%v, Len(%d)\n", nums, len(nums))

}

func numbertestprime(num int) {

	if prime.PrimeNumber(num) {

		mu.Lock()

		nums = append(nums, num)

		mu.Unlock()
	}

}
