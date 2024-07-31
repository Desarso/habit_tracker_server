package controllers

import (
	"github.com/desarso/habit_tracker_server/helpers/DBManager"
	"github.com/desarso/habit_tracker_server/models"
	"github.com/gin-gonic/gin"
)

type Habit_Controller struct{}

// @Summary Add a habit
// @Description Add a habit with name and description
// @Tags habits
// @Accept  json
// @Produce  json
// @Param name query string true "Habit name"
// @Param description query string true "Habit description"
// @Success 200 {string} string "Succesfully added habit"
// @Failure 400 {string} string "Both 'name' and 'description' are required"
// @Router /habits/addHabit [post]
func (h *Habit_Controller) AddHabit(c *gin.Context) {
	name := c.Query("name")
	description := c.Query("description")

	if name == "" || description == "" {
		c.JSON(400, gin.H{"error": "Both 'name' and 'description' are required"})
		return
	}

	// Create a habit object
	habit := models.Habit{
		Name:        name,
		Description: description,
	}

	//save habit to database in habits collection
	_, err := DBManager.DB.Collection("habits").InsertOne(c, habit)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Succesfully added habit"})
}

// @Summary Get all habits
// @Description Get all habits from the database
// @Tags habits
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.Habit "List of habits"
// @Failure 500 {string} string "Internal server error"
// @Router /habits/getHabits [get]
func (h *Habit_Controller) GetHabits(c *gin.Context) {
	//get all habits from the database
	cursor, err := DBManager.DB.Collection("habits").Find(c, gin.H{})
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	var habits []models.Habit
	if err = cursor.All(c, &habits); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, habits)
}
