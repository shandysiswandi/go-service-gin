package main

import (
	"go-service-gin/application"
	"go-service-gin/config"
	"go-service-gin/infrastructure/database"
	"go-service-gin/infrastructure/database/mysql"
	"go-service-gin/interfaces/console"
	"go-service-gin/util/logger"
	"sync"
	"time"

	"github.com/joho/godotenv"
)

var wg = sync.WaitGroup{}

func main() {
	/********** ********** ********** ********** ********** ********** ********** ********** ********** **********/
	/* load environment variable
	/********** ********** ********** ********** ********** ********** ********** ********** ********** **********/
	if err := godotenv.Load(); err != nil {
		logger.Error(err)
	}

	/********** ********** ********** ********** ********** ********** ********** ********** ********** **********/
	/* define configuration
	/********** ********** ********** ********** ********** ********** ********** ********** ********** **********/
	var (
		dbConfig = config.NewDatabaseConfig()
	)

	/********** ********** ********** ********** ********** ********** ********** ********** ********** **********/
	/* define database connection
	/********** ********** ********** ********** ********** ********** ********** ********** ********** **********/
	db, err := database.NewDatabase(dbConfig)
	if err != nil {
		logger.Error(err)
	}

	/********** ********** ********** ********** ********** ********** ********** ********** ********** **********/
	/* define variable and inject to constructor
	/********** ********** ********** ********** ********** ********** ********** ********** ********** **********/
	var (
		blogRepository = mysql.NewBlogRepository(db)
		blogLogic      = application.NewBlogApplication(blogRepository)
		blogConsole    = console.NewBlogConsole(blogLogic)
	)

	// testing. stress test
	t := time.Now()
	for i := 1; i <= 5; i++ {
		blogConsole.Create()
	}
	logger.Info(time.Since(t))
}
