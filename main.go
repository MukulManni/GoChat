package main

import (
	"os"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func main() {

	port := os.Getenv("PORT")
	r = gin.Default()
	r.LoadHTMLGlob("static/template/*")

	r.Use(sessions.Sessions("mysession", sessions.NewCookieStore([]byte("secret"))))

	intializeRoutes()

	r.Run(":" + port)
}
