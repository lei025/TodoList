package routes

import (
	"todo-list/api"
	"todo-list/middleware"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	store := cookie.NewStore([]byte("somthing-very-secret"))
	r.Use(sessions.Sessions("mysession", store))

	v1 := r.Group("api/v1")
	{
		// TODO: 用户操作
		v1.POST("user/register", api.UserRegister)
		v1.POST("user/login", api.UserLogin)

		authed := v1.Group("/")
		authed.Use(middleware.JWT())
		{
			authed.POST("task", api.CreateTask)
		}
	}

	return r
}
