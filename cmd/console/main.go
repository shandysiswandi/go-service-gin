package main

import (
	app "go-service-gin"
	"go-service-gin/application"
	"go-service-gin/infrastructure/external/jsonplaceholder"
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
		jph       = appData[app.JSONPlaceholder].(*jsonplaceholder.JPH)

		blogConsole = console.NewBlogConsole(blogLogic)
	)

	// testing. stress test
	t := time.Now()
	for i := 1; i <= 1; i++ {
		blogConsole.Create()
	}
	logger.Info(time.Since(t))

	aa, ee := jph.Posts.Fetch()
	logger.Error(ee)
	logger.Info(aa)
}
