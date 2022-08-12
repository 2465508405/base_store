package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	c.HTML(http.StatusOK, "login/login.html", gin.H{"title": "后台管理系统", "address": "www.5lmh.com"})

}

func AuthLogin(c *gin.Context) {

}
