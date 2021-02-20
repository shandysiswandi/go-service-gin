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
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	/********** ********** ********** ********** ********** ********** ********** ********** ********** **********/
	/* load environment variable
	/********** ********** ********** ********** ********** ********** ********** ********** ********** **********/
	if err := godotenv.Load(); err != nil {
		logger.LogError(err)
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
		logger.LogError(err)
	}

	/********** ********** ********** ********** ********** ********** ********** ********** ********** **********/
	/* define library
	/********** ********** ********** ********** ********** ********** ********** ********** ********** **********/
	redis := redis.New(redisConfig)
	sentry := sentry.New(sentryConfig)

	logger.LogInfo(redis)
	logger.LogInfo(sentry)

	/********** ********** ********** ********** ********** ********** ********** ********** ********** **********/
	/* define variable and inject to constructor
	/********** ********** ********** ********** ********** ********** ********** ********** ********** **********/
	var (
		// default
		defaultHandler = rest.NewDefaultREST()

		// blogs
		blogRepository = mysql.NewBlogRepository(db)
		blogLogic      = application.NewBlogApplication(blogRepository)
		blogHandler    = rest.NewBlogREST(blogLogic)
	)

	/********** ********** ********** ********** ********** ********** ********** ********** ********** **********/
	/* define engine instance from gin framework
	/********** ********** ********** ********** ********** ********** ********** ********** ********** **********/
	engine := gin.Default()
	if appConfig.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	/********** ********** ********** ********** ********** ********** ********** ********** ********** **********/
	/* define gin middleware
	/********** ********** ********** ********** ********** ********** ********** ********** ********** **********/
	engine.Use(gzip.Gzip(gzip.DefaultCompression))
	engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           2 * time.Hour,
	}))
	engine.Use(func(c *gin.Context) {
		c.Writer.Header().Set("X-Content-Type-Options", "nosniff")
		c.Writer.Header().Set("X-DNS-Prefetch-Control", "off")
		c.Writer.Header().Set("X-Frame-Options", "DENY")
		c.Writer.Header().Set("Strict-Transport-Security", "true; includeSubDomains")
		c.Writer.Header().Set("X-Download-Options", "noopen")
		c.Writer.Header().Set("X-XSS-Protection", "1; mode=block")
		c.Writer.Header().Set("Referrer-Policy", "no-referrer")
	})

	/********** ********** ********** ********** ********** ********** ********** ********** ********** **********/
	/* define routes
	/********** ********** ********** ********** ********** ********** ********** ********** ********** **********/
	engine.StaticFile("/favicon.ico", "./resource/media/favicon.ico")
	engine.NoRoute(defaultHandler.RouteNotFound)
	engine.HandleMethodNotAllowed = true
	engine.NoMethod(defaultHandler.MethodNotFound)
	engine.GET("/", defaultHandler.Home)

	br := engine.Group("/blogs")
	{
		if blogHandler != nil {
			br.GET("", blogHandler.Fetch)
		}
	}

	/********** ********** ********** ********** ********** ********** ********** ********** ********** **********/
	/* running server using engine instance from gin framework
	/********** ********** ********** ********** ********** ********** ********** ********** ********** **********/
	engine.Run(fmt.Sprintf("%s:%s", appConfig.Host, appConfig.Port))
}
