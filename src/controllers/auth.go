package controllers

import (
	"net/http"

	"theater-production/src/services"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func GetSessionUser(c *gin.Context) gin.H {
	session := sessions.Default(c)
	return gin.H{
		"UserID":   session.Get("user_id"),
		"Username": session.Get("username"),
		"Role":     session.Get("role"),
	}
}

func LoginHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "auth/login.html", gin.H{
		"Title": "登录",
	})
}

func LoginPost(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	user, err := services.AuthenticateUser(username, password)
	if err != nil {
		c.HTML(http.StatusUnauthorized, "auth/login.html", gin.H{
			"Title": "登录",
			"Error": "用户名或密码错误",
		})
		return
	}

	session := sessions.Default(c)
	session.Set("user_id", user.ID)
	session.Set("username", user.Username)
	session.Set("role", user.Role)
	session.Save()

	c.Redirect(http.StatusFound, "/dashboard")
}

func RegisterHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "auth/register.html", gin.H{
		"Title": "注册",
	})
}

func RegisterPost(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	confirmPassword := c.PostForm("confirm_password")
	role := c.PostForm("role")

	if password != confirmPassword {
		c.HTML(http.StatusBadRequest, "auth/register.html", gin.H{
			"Title": "注册",
			"Error": "两次输入的密码不一致",
		})
		return
	}

	if err := services.CreateUser(username, password, role); err != nil {
		c.HTML(http.StatusBadRequest, "auth/register.html", gin.H{
			"Title": "注册",
			"Error": "注册失败，用户名可能已存在",
		})
		return
	}

	c.Redirect(http.StatusFound, "/auth/login")
}

func LogoutHandler(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.Redirect(http.StatusFound, "/auth/login")
}
