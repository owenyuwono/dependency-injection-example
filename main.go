package main

import (
	"database/sql"
	"dependency-injection-example/handler"
	"dependency-injection-example/repository"
	"log"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

func main() {
	router := gin.Default()
	db, err := sql.Open("postgres", "postgres://owenyuwono:root@localhost:5432/exampledb?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := repository.New(db)
	h := handler.New(r)

	router.GET("/healthcheck", h.Healthcheck)
	router.POST("/insert", h.InsertData)

	router.Run()
}
