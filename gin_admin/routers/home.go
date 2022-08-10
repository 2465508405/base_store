package routers

import (
	"project/allfunc/gin_admin/api"

	"github.com/gin-gonic/gin"
)

func InitHome(r *gin.Engine) {

	r.GET("/home", api.Home)
}
