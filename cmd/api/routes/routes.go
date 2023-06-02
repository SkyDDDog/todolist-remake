package routes

import (
	"github.com/gin-gonic/gin"
	"todolist-remake/cmd/api/handler"
	"todolist-remake/cmd/api/middleware"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", handler.Ping)

	user := r.Group("/user")
	{
		user.POST("/register", handler.UserRegister)
		user.POST("/login", handler.UserLogin)
	}

	task := r.Group("/task").Use(middleware.JWT)
	{
		task.POST("", handler.TaskCreate)
		task.GET("", handler.TaskList)
		task.PUT("", handler.TaskUpdate)
		task.DELETE("", handler.TaskDelete)
		task.POST("/done", handler.TaskDone)
		task.POST("/undo", handler.TaskUnDone)
	}

	return r
}
