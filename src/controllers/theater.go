package controllers

import (
	"net/http"
	"strconv"
	"time"

	"theater-production/src/models"
	"theater-production/src/services"

	"github.com/gin-gonic/gin"
)

func TheaterList(c *gin.Context) {
	theaters, err := services.GetAllTheaters()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "theaters/list.html", gin.H{
			"Title": "剧场档期",
			"Error": "获取剧场列表失败",
		})
		return
	}
	user := GetSessionUser(c)
	c.HTML(http.StatusOK, "theaters/list.html", gin.H{
		"Title":    "剧场档期",
		"Theaters": theaters,
		"User":     user,
	})
}

func TheaterDetail(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.Redirect(http.StatusFound, "/theaters/")
		return
	}
	theater, err := services.GetTheaterByID(uint(id))
	if err != nil {
		c.Redirect(http.StatusFound, "/theaters/")
		return
	}
	schedules, _ := services.GetTheaterSchedulesByTheaterID(uint(id))
	user := GetSessionUser(c)
	c.HTML(http.StatusOK, "theaters/detail.html", gin.H{
		"Title":     theater.Name,
		"Theater":   theater,
		"Schedules": schedules,
		"User":      user,
	})
}

func TheaterScheduleView(c *gin.Context) {
	schedules, err := services.GetAllTheaterSchedules()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "theaters/schedule.html", gin.H{
			"Title": "档期管理",
			"Error": "获取档期信息失败",
		})
		return
	}
	theaters, _ := services.GetAllTheaters()
	scripts, _ := services.GetAllScripts()
	user := GetSessionUser(c)
	c.HTML(http.StatusOK, "theaters/schedule.html", gin.H{
		"Title":     "档期管理",
		"Schedules": schedules,
		"Theaters":  theaters,
		"Scripts":   scripts,
		"User":      user,
	})
}

func TheaterSchedulePost(c *gin.Context) {
	theaterID, _ := strconv.ParseUint(c.PostForm("theater_id"), 10, 32)
	scriptID, _ := strconv.ParseUint(c.PostForm("script_id"), 10, 32)
	startTime, err := time.Parse("2006-01-02T15:04", c.PostForm("start_time"))
	if err != nil {
		startTime = time.Now()
	}
	endTime, err := time.Parse("2006-01-02T15:04", c.PostForm("end_time"))
	if err != nil {
		endTime = time.Now().Add(2 * time.Hour)
	}
	schedule := &models.TheaterSchedule{
		TheaterID: uint(theaterID),
		ScriptID:  uint(scriptID),
		StartTime: startTime,
		EndTime:   endTime,
		Status:    c.PostForm("status"),
	}
	if schedule.Status == "" {
		schedule.Status = "booked"
	}
	if err := services.CreateTheaterSchedule(schedule); err != nil {
		c.Redirect(http.StatusFound, "/theaters/schedule")
		return
	}
	c.Redirect(http.StatusFound, "/theaters/schedule")
}
