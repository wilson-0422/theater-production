package controllers

import (
	"net/http"
	"strconv"
	"time"

	"theater-production/src/models"
	"theater-production/src/services"

	"github.com/gin-gonic/gin"
)

func TourList(c *gin.Context) {
	tours, err := services.GetAllTours()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "tours/list.html", gin.H{
			"Title": "巡演行程",
			"Error": "获取巡演列表失败",
		})
		return
	}
	user := GetSessionUser(c)
	c.HTML(http.StatusOK, "tours/list.html", gin.H{
		"Title": "巡演行程",
		"Tours": tours,
		"User":  user,
	})
}

func TourDetail(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.Redirect(http.StatusFound, "/tours/")
		return
	}
	tour, err := services.GetTourByID(uint(id))
	if err != nil {
		c.Redirect(http.StatusFound, "/tours/")
		return
	}
	user := GetSessionUser(c)
	c.HTML(http.StatusOK, "tours/detail.html", gin.H{
		"Title": "巡演详情",
		"Tour":  tour,
		"User":  user,
	})
}

func TourCreate(c *gin.Context) {
	scripts, _ := services.GetAllScripts()
	user := GetSessionUser(c)
	c.HTML(http.StatusOK, "tours/create.html", gin.H{
		"Title":   "新建巡演",
		"Scripts": scripts,
		"User":    user,
	})
}

func TourCreatePost(c *gin.Context) {
	scriptID, _ := strconv.ParseUint(c.PostForm("script_id"), 10, 32)
	startDate, err := time.Parse("2006-01-02", c.PostForm("start_date"))
	if err != nil {
		startDate = time.Now()
	}
	endDate, err := time.Parse("2006-01-02", c.PostForm("end_date"))
	if err != nil {
		endDate = time.Now().Add(7 * 24 * time.Hour)
	}
	tour := &models.Tour{
		ScriptID:  uint(scriptID),
		City:      c.PostForm("city"),
		Venue:     c.PostForm("venue"),
		StartDate: startDate,
		EndDate:   endDate,
		Status:    c.PostForm("status"),
		Notes:     c.PostForm("notes"),
	}
	if tour.Status == "" {
		tour.Status = "planned"
	}
	if err := services.CreateTour(tour); err != nil {
		scripts, _ := services.GetAllScripts()
		user := GetSessionUser(c)
		c.HTML(http.StatusBadRequest, "tours/create.html", gin.H{
			"Title":   "新建巡演",
			"Error":   "创建巡演失败",
			"Scripts": scripts,
			"User":    user,
		})
		return
	}
	c.Redirect(http.StatusFound, "/tours/")
}
