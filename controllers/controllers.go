package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todoAPI/models"
)

type NewTask struct {
	Name string `json:"name" binding:"required"`
}

type UpdateTask struct {
	IsCompleted bool `json:"isCompleted"`
}

func GetAllTodos(c *gin.Context) {
	var todos []models.TodoUnit
	models.DB.Debug().Find(&todos)

	c.JSON(http.StatusOK, todos)
}

func GetTask(c *gin.Context) {
	var task models.TodoUnit
	if err := models.DB.Debug().Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, task)
}

func CreateTask(c *gin.Context) {
	var input NewTask
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	task := models.TodoUnit{Name: input.Name}
	models.DB.Debug().Create(&task)
	c.JSON(http.StatusOK, task)
}

func CompleteTask(c *gin.Context) {
	var task models.TodoUnit
	if err := models.DB.Debug().Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var isDone UpdateTask
	if err := c.ShouldBindJSON(&isDone); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.DB.Debug().Model(&task).Update("is_completed", isDone.IsCompleted)

	c.JSON(http.StatusOK, task)
}
