package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine
var db *sql.DB

func main() {
	url, ok := os.LookupEnv("DATABASE_URL")

	if !ok {
		log.Fatalln("Database url is required.")
	}

	var err error
	db, err = connectDB(url)

	if err != nil {
		log.Fatalf("Error connecting database: %s", err.Error())
	}

	port := os.Getenv("PORT")

	msglist, err := getAllMsgsDB(db)

	if err != nil {
		log.Fatalln("Unable to retrieve messages from database.")
	}

	globalmsgList = append(globalmsgList, msglist...)

	r = gin.Default()
	r.LoadHTMLGlob("static/template/*")

	r.Use(sessions.Sessions("chatsession", sessions.NewCookieStore([]byte("secret"))))

	intializeRoutes()

	r.Run(":" + port)
}
