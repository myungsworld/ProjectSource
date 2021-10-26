package main

import (
	"ProjectSource/backend/config"
	docs "ProjectSource/backend/docs"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	osSignal chan os.Signal
)

const PORT = 5011

func main() {

	osSignal = make(chan os.Signal, 10000)
	signal.Notify(osSignal, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// 설정값 초기화
	config.Init(config.Dev)

	logFlags := log.Lshortfile
	logFlags = logFlags | log.Ldate | log.Ltime
	log.SetFlags(logFlags)

	r := gin.New()
	r.Use(gin.Logger())

	// cors
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete},
		AllowHeaders: []string{"Accept", "Accept-Language", "Content-Language", "Origin", "Content-Type", "Authorization"},
	}))

	r.Use(gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		return fmt.Sprintf("%d | %v | %s | %s\n", params.StatusCode, params.Latency, params.Method, params.Path)
	}))

	srv := &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf(":%d", PORT),
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	// 스웨거 BasePath 설정
	docs.SwaggerInfo.BasePath = ""

	rApi := r.Group("/api")
	//rApi.Use(middlewares.WrapMiddleware)

	rTrade := rApi.Group("/trade")
	SetWebTrade(rTrade)

	servingFE(r)

	go func() {
		log.Println("Starting Server")
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	<-osSignal

	fmt.Println("Terminating server")

}

func servingFE(r *gin.Engine) {
	r.NoRoute(func(c *gin.Context) {
		c.File("../frontend/src/app/app.component.index.html")
	})
}
