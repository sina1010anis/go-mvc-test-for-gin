package main

import (
	"fmt"
	"sync"
	"web/local_routes"
	// "github.com/gin-gonic/gin"
)

func main() {

	local_routes.BuildRoute()

}

func pushJob(job chan int, num int, wg *sync.WaitGroup) {
	defer wg.Done()
	job <- num
}

func readJob(job chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range job {

		fmt.Printf("Jon => %d----", val)

	}
}
