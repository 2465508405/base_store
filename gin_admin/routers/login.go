package routers

import (
	"project/allfunc/gin_admin/api"

	"github.com/gin-gonic/gin"
)

func LoginRouter(r *gin.Engine) {

	r.GET("/login", api.Login)
	r.POST("/auth/login", api.AuthLogin)

}
