package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
	"github.com/joho/godotenv"

	database "course-chart/config"
)

type Course struct {
	Id          int
	Name        string
	Institution string
	CreditHours int
	Length      int
	CreatedAt   string
	UpdatedAt   string
}

func setupRouter(db *pg.DB) *gin.Engine {
	r := gin.Default()
	r.GET("/courses/1", func(c *gin.Context) {

		course := &Course{Id: 1}
		err := db.Model(course).WherePK().Select()
		if err != nil {
			panic(err)
		}

		c.String(200, course.Name)
	})
	return r
}

func main() {
	godotenv.Load()

	port := ":" + os.Getenv("PORT")

	db := database.Connect()

	router := setupRouter(db)

	log.Fatal(router.Run(port))
}
