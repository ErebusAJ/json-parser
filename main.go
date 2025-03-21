package main

import (
	"log"
	"os"
	"time"

	"github.com/ErebusAJ/bajajAPI/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main(){
	godotenv.Load()

	port := os.Getenv("PORT_NO")
	if port == ""{
		log.Printf("error retreiving port")
	}
	//initialize router
	r := gin.Default()

	r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"*"},  // In production, replace with your frontend domain
        AllowMethods:     []string{"POST", "GET"},
        AllowHeaders:     []string{"Origin", "Content-Type"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge:          12 * time.Hour,
    }))

	r.Use(func(c *gin.Context) {
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("X-Frame-Options", "DENY")
		c.Header("X-XSS-Protection", "1; mode=block")
		c.Next()
	})

	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context){
		c.HTML(200, "index.html", nil)
	})
	r.POST("/bfhl", utils.HandleArray)
	r.GET("/bfhl", utils.HandleOperation)

	r.Run(":"+port)
}