package controllers

import (
	"github.com/desarso/habit_tracker_server/models"
	"github.com/gin-gonic/gin"
)

type Habit_Controller struct{}

// AddHabit godoc
// @Summary Add a habit
// @Description Add a habit
// @Tags habits
// @Accept  json
// @Produce  json
// @Param habit body Habit true "Habit object"
// @Success 200 {object} Habit
// @Router /addHabit [post]
func (h *Habit_Controller) AddHabit(c *gin.Context) {
	var habit models.Habit
	c.BindJSON(&habit)

}
