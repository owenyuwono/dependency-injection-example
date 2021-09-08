package handler

import (
	"log"
	"net/http"

	"dependency-injection-example/model"

	"github.com/gin-gonic/gin"
)

//go:generate mockgen -source=./handler.go -destination=./mock/mock.go -package=mock
type repository interface {
	Insert(data model.Data) (string, error)
}

type Handler struct {
	repo repository
}

func New(repo repository) *Handler {
	return &Handler{
		repo: repo,
	}
}

func (h *Handler) Healthcheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}

func (h *Handler) InsertData(c *gin.Context) {
	var body model.Data
	err := c.BindJSON(&body)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "bad request",
		})
		return
	}
	// add some new requirements
	id, err := h.repo.Insert(body)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "internal server error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"id":     id,
	})
}
