package app

import (
	"go-service-gin/application"
	"go-service-gin/config"
	"go-service-gin/infrastructure/database"
	"go-service-gin/infrastructure/database/mysql"
	"go-service-gin/infrastructure/library/redis"
	"go-service-gin/infrastructure/library/sentry"
	"go-service-gin/util/logger"

	"github.com/joho/godotenv"
)

// all key app data
const (
	AppConfig = iota
	Sentry    = iota
	Redis     = iota
	BlogLogic = iota
)

// New is
func New() map[int]interface{} {
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
		appConfig    = config.NewAppConfig()
		dbConfig     = config.NewDatabaseConfig()
		redisConfig  = config.NewRedisConfig()
		sentryConfig = config.NewSentryConfig()
	)

	/********** ********** ********** ********** ********** ********** ********** ********** ********** **********/
	/* define database connection
	/********** ********** ********** ********** ********** ********** ********** ********** ********** **********/
	db, err := database.NewDatabase(dbConfig)
	if err != nil {
		logger.Error(err)
	}

	/********** ********** ********** ********** ********** ********** ********** ********** ********** **********/
	/* define library
	/********** ********** ********** ********** ********** ********** ********** ********** ********** **********/
	redis := redis.New(redisConfig)
	sentry := sentry.New(sentryConfig)

	/********** ********** ********** ********** ********** ********** ********** ********** ********** **********/
	/* define variable and inject to constructor
	/********** ********** ********** ********** ********** ********** ********** ********** ********** **********/
	var (
		// blogs
		blogRepository = mysql.NewBlogRepository(db)
		blogLogic      = application.NewBlogApplication(blogRepository)
	)

	/********** ********** ********** ********** ********** ********** ********** ********** ********** **********/
	/* return data
	/********** ********** ********** ********** ********** ********** ********** ********** ********** **********/
	return map[int]interface{}{
		AppConfig: appConfig,
		Sentry:    sentry,
		Redis:     redis,
		BlogLogic: blogLogic,
	}
}
