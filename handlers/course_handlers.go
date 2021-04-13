package handlers

import (
	db "course-chart/config"
	"course-chart/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RenderError(c *gin.Context, err error) {
	log.Printf("Error retrieving course from database.\nReason: %v", err)
	c.JSON(http.StatusNotFound, gin.H{
		"status": http.StatusNotFound,
		"errors": "Course not found",
	})
}

func GetCourse(c *gin.Context) {
	var course models.Course
	err := db.Conn.Preload("Modules.ModuleActivities.Activity").First(&course, c.Param("id")).Error

	if err != nil {
		RenderError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Course found",
		"data":    course,
	})
}
