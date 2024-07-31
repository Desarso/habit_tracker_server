package routes

import (
	"github.com/desarso/habit_tracker_server/controllers"
	"github.com/gin-gonic/gin"
)

var Habit_Controllers = controllers.Habit_Controller{}

func Habit_Routes(r *gin.RouterGroup) {
	r.POST("/addHabit", Habit_Controllers.AddHabit)
	r.GET("/getHabits", Habit_Controllers.GetHabits)
}
