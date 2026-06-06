package controllers

import (
	"net/http"
	"strconv"

	"theater-production/src/models"
	"theater-production/src/services"

	"github.com/gin-gonic/gin"
)

func ScriptList(c *gin.Context) {
	scripts, err := services.GetAllScripts()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "scripts/list.html", gin.H{
			"Title": "剧本存档",
			"Error": "获取剧本列表失败",
		})
		return
	}
	user := GetSessionUser(c)
	c.HTML(http.StatusOK, "scripts/list.html", gin.H{
		"Title":   "剧本存档",
		"Scripts": scripts,
		"User":    user,
	})
}

func ScriptDetail(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.Redirect(http.StatusFound, "/scripts/")
		return
	}
	script, err := services.GetScriptByID(uint(id))
	if err != nil {
		c.Redirect(http.StatusFound, "/scripts/")
		return
	}
	user := GetSessionUser(c)
	c.HTML(http.StatusOK, "scripts/detail.html", gin.H{
		"Title":  script.Title,
		"Script": script,
		"User":   user,
	})
}

func ScriptCreate(c *gin.Context) {
	user := GetSessionUser(c)
	c.HTML(http.StatusOK, "scripts/create.html", gin.H{
		"Title": "新建剧本",
		"User":  user,
	})
}

func ScriptCreatePost(c *gin.Context) {
	duration, _ := strconv.Atoi(c.PostForm("duration"))
	script := &models.Script{
		Title:    c.PostForm("title"),
		Author:   c.PostForm("author"),
		Genre:    c.PostForm("genre"),
		Synopsis: c.PostForm("synopsis"),
		Duration: duration,
		Status:   c.PostForm("status"),
	}
	if script.Status == "" {
		script.Status = "draft"
	}
	if err := services.CreateScript(script); err != nil {
		user := GetSessionUser(c)
		c.HTML(http.StatusBadRequest, "scripts/create.html", gin.H{
			"Title": "新建剧本",
			"Error": "创建剧本失败",
			"User":  user,
		})
		return
	}
	c.Redirect(http.StatusFound, "/scripts/")
}

func ScriptEdit(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.Redirect(http.StatusFound, "/scripts/")
		return
	}
	script, err := services.GetScriptByID(uint(id))
	if err != nil {
		c.Redirect(http.StatusFound, "/scripts/")
		return
	}
	user := GetSessionUser(c)
	c.HTML(http.StatusOK, "scripts/edit.html", gin.H{
		"Title":  "编辑剧本",
		"Script": script,
		"User":   user,
	})
}

func ScriptEditPost(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.Redirect(http.StatusFound, "/scripts/")
		return
	}
	script, err := services.GetScriptByID(uint(id))
	if err != nil {
		c.Redirect(http.StatusFound, "/scripts/")
		return
	}
	duration, _ := strconv.Atoi(c.PostForm("duration"))
	script.Title = c.PostForm("title")
	script.Author = c.PostForm("author")
	script.Genre = c.PostForm("genre")
	script.Synopsis = c.PostForm("synopsis")
	script.Duration = duration
	script.Status = c.PostForm("status")
	if err := services.UpdateScript(script); err != nil {
		user := GetSessionUser(c)
		c.HTML(http.StatusBadRequest, "scripts/edit.html", gin.H{
			"Title":  "编辑剧本",
			"Error":  "更新剧本失败",
			"Script": script,
			"User":   user,
		})
		return
	}
	c.Redirect(http.StatusFound, "/scripts/"+c.Param("id"))
}

func ScriptDelete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.Redirect(http.StatusFound, "/scripts/")
		return
	}
	services.DeleteScript(uint(id))
	c.Redirect(http.StatusFound, "/scripts/")
}
