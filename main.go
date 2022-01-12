package main

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func main() {

	//port := os.Getenv("PORT")
	r = gin.Default()
	r.LoadHTMLGlob("static/template/*")

	r.Use(sessions.Sessions("chatsession", sessions.NewCookieStore([]byte("secret"))))

	intializeRoutes()

	r.Run(":8080")
}
