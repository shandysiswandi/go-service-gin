package main

import (
	"fmt"
	"go-service-gin/application"
	"go-service-gin/config"
	"go-service-gin/infrastructure/database"
	"go-service-gin/infrastructure/database/mysql"
	"go-service-gin/infrastructure/library/redis"
	"go-service-gin/infrastructure/library/sentry"
	"go-service-gin/interfaces/rest"
	"go-service-gin/util/logger"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// load environment variable
	if err := godotenv.Load(); err != nil {
		logger.LogError(err)
	}

	// define configuration
	var (
		appConfig    = config.NewAppConfig()
		dbConfig     = config.NewDatabaseConfig()
		redisConfig  = config.NewRedisConfig()
		sentryConfig = config.NewSentryConfig()
	)

	// define database connection
	db, err := database.NewDatabase(dbConfig)
	if err != nil {
		logger.LogError(err)
	}

	// define library
	redis := redis.New(redisConfig)
	sentry := sentry.New(sentryConfig)

	logger.LogInfo(redis)
	logger.LogInfo(sentry)

	// define variable and inject to constructor
	var (
		// home
		homeHandler = rest.NewHomeREST()

		// blogs
		blogRepository = mysql.NewBlogRepository(db)
		blogLogic      = application.NewBlogApplication(blogRepository)
		blogHandler    = rest.NewBlogREST(blogLogic)
	)

	// define engine instance from gin framework
	engine := gin.Default()

	// define router
	engine.GET("/", homeHandler.Home)

	// router /blogs
	blogRouter := engine.Group("/blogs")
	blogRouter.GET("", blogHandler.Fetch)

	// running server using engine instance from gin framework
	engine.Run(fmt.Sprintf("%s:%s", appConfig.Host, appConfig.Port))
}
