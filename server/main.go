package main


import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"server/router"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	router.RegisterRouter(r)
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
