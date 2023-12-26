package router

import (
	"github.com/gin-gonic/gin"
	"server/handler"
)

func RegisterRouter(r *gin.Engine) {


	r.POST("/login", handler.LoginHandler)
	r.POST("/register", handler.RegisterHandler)

	r.GET("/user", handler.GetUser)
	r.POST("/user", handler.AddUser)
	r.PUT("/user", handler.UpdateUser)
	r.DELETE("/user", handler.DeleteUser)

	r.GET("/role", handler.GetRole)
	r.GET("/roleaccess", handler.GetAccessByRoleIDHandler)
	r.POST("/role", handler.AddRole)
	r.PUT("/role", handler.UpdateRole)
	r.DELETE("/role:id", handler.DeleteRole)

	r.GET("/access", handler.GetAccess)
	r.POST("/access", handler.AddAccess)
	r.PUT("/access", handler.UpdateAccess)
	r.DELETE("/access", handler.DeleteAccess)

}