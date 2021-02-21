package main

import (
	app "go-service-gin"
	"go-service-gin/application"
	"go-service-gin/interfaces/console"
	"go-service-gin/util/logger"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {
	/* load app data and typed and define variable hanlder and inject data to constructor */
	var (
		appData   = app.New()
		blogLogic = appData[app.BlogLogic].(*application.BlogApplication)

		blogConsole = console.NewBlogConsole(blogLogic)
	)

	// testing. stress test
	t := time.Now()
	for i := 1; i <= 5; i++ {
		blogConsole.Create()
	}
	logger.Info(time.Since(t))
}
