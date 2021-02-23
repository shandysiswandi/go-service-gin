package app

import (
	"go-service-gin/application"
	"go-service-gin/config"
	"go-service-gin/infrastructure/database"
	"go-service-gin/infrastructure/database/sql"
	"go-service-gin/infrastructure/external/jsonplaceholder"
	"go-service-gin/infrastructure/library/redis"
	"go-service-gin/infrastructure/library/sentry"
	"go-service-gin/util/logger"

	"github.com/joho/godotenv"
)

// all key app data
const (
	AppConfig       = iota
	Sentry          = iota
	Redis           = iota
	JSONPlaceholder = iota
	BlogLogic       = iota
)

// New is
func New() map[int]interface{} {
	/********** ********** ********** ********** ********** ********** ********** ********** ********** **********/
	/* load environment variable & define configuration
	/********** ********** ********** ********** ********** ********** ********** ********** ********** **********/
	if err := godotenv.Load(); err != nil {
		logger.Error(err)
	}

	var (
		appConfig      = config.NewAppConfig()
		dbConfig       = config.NewDatabaseConfig()
		redisConfig    = config.NewRedisConfig()
		sentryConfig   = config.NewSentryConfig()
		externalConfig = config.NewExternalConfig()
	)

	/********** ********** ********** ********** ********** ********** ********** ********** ********** **********/
	/* define infrasuctures | database | library | external |
	/********** ********** ********** ********** ********** ********** ********** ********** ********** **********/
	db, err := database.NewDatabase(dbConfig)
	if err != nil {
		logger.Error(err)
	}

	redis := redis.New(redisConfig)
	sentry := sentry.New(sentryConfig)

	jph := jsonplaceholder.New(externalConfig.JSONPlaceholder)

	/********** ********** ********** ********** ********** ********** ********** ********** ********** **********/
	/* define variable and inject to constructor
	/********** ********** ********** ********** ********** ********** ********** ********** ********** **********/
	var (
		// blogs
		blogRepository = sql.NewBlogRepository(db)
		blogLogic      = application.NewBlogApplication(blogRepository)
	)

	/********** ********** ********** ********** ********** ********** ********** ********** ********** **********/
	/* return data
	/********** ********** ********** ********** ********** ********** ********** ********** ********** **********/
	return map[int]interface{}{
		AppConfig:       appConfig,
		Sentry:          sentry,
		Redis:           redis,
		JSONPlaceholder: jph,
		BlogLogic:       blogLogic,
	}
}
