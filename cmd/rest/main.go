package main

import (
	"fmt"
	app "go-service-gin"
	"go-service-gin/application"
	"go-service-gin/config"
	"go-service-gin/infrastructure/library/sentry"
	"go-service-gin/interfaces/rest"
	"go-service-gin/util/logger"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func main() {
	/* load app data and typed and define variable hanlder and inject data to constructor */
	var (
		appData   = app.New()
		appConfig = appData[app.AppConfig].(*config.AppConfig)
		sentry    = appData[app.Sentry].(*sentry.Sentry)
		blogLogic = appData[app.BlogLogic].(*application.BlogApplication)

		defaultHandler = rest.NewDefaultREST(sentry)
		blogHandler    = rest.NewBlogREST(blogLogic)
	)

	/* define engine instance from gin framework */
	engine := gin.Default()
	if appConfig.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	middlewares(engine)

	/* define routes */
	engine.StaticFile("/favicon.ico", "./resource/media/favicon.ico")
	engine.NoRoute(defaultHandler.RouteNotFound)
	engine.HandleMethodNotAllowed = true
	engine.NoMethod(defaultHandler.MethodNotFound)
	engine.GET("/", defaultHandler.Home)
	engine.GET("/check/sentry", defaultHandler.Sentry)

	if blogHandler != nil {
		engine.GET("/blogs", blogHandler.Fetch)
	}

	/* running server using engine instance from gin framework */
	if err := engine.Run(fmt.Sprintf("%s:%s", appConfig.Host, appConfig.Port)); err != nil {
		logger.Error(err)
	}
}

func middlewares(e *gin.Engine) {
	e.Use(gzip.Gzip(gzip.DefaultCompression))

	e.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           2 * time.Hour,
	}))

	e.Use(func(c *gin.Context) {
		c.Writer.Header().Set("X-Content-Type-Options", "nosniff")
		c.Writer.Header().Set("X-DNS-Prefetch-Control", "off")
		c.Writer.Header().Set("X-Frame-Options", "DENY")
		c.Writer.Header().Set("Strict-Transport-Security", "true; includeSubDomains")
		c.Writer.Header().Set("X-Download-Options", "noopen")
		c.Writer.Header().Set("X-XSS-Protection", "1; mode=block")
		c.Writer.Header().Set("Referrer-Policy", "no-referrer")
	})
}
