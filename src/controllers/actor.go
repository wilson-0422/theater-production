package controllers

import (
	"net/http"
	"strconv"
	"time"

	"theater-production/src/models"
	"theater-production/src/services"

	"github.com/gin-gonic/gin"
)

func ActorList(c *gin.Context) {
	actors, err := services.GetAllActors()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "actors/list.html", gin.H{
			"Title": "演员排班",
			"Error": "获取演员列表失败",
		})
		return
	}
	user := GetSessionUser(c)
	c.HTML(http.StatusOK, "actors/list.html", gin.H{
		"Title":  "演员排班",
		"Actors": actors,
		"User":   user,
	})
}

func ActorDetail(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.Redirect(http.StatusFound, "/actors/")
		return
	}
	actor, err := services.GetActorByID(uint(id))
	if err != nil {
		c.Redirect(http.StatusFound, "/actors/")
		return
	}
	schedules, _ := services.GetActorSchedules(uint(id))
	user := GetSessionUser(c)
	c.HTML(http.StatusOK, "actors/detail.html", gin.H{
		"Title":      actor.Name,
		"Actor":      actor,
		"Schedules":  schedules,
		"User":       user,
	})
}

func ActorScheduleView(c *gin.Context) {
	schedules, err := services.GetAllActorSchedules()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "actors/schedule.html", gin.H{
			"Title": "排班管理",
			"Error": "获取排班信息失败",
		})
		return
	}
	actors, _ := services.GetAllActors()
	scripts, _ := services.GetAllScripts()
	user := GetSessionUser(c)
	c.HTML(http.StatusOK, "actors/schedule.html", gin.H{
		"Title":      "排班管理",
		"Schedules":  schedules,
		"Actors":     actors,
		"Scripts":    scripts,
		"User":       user,
	})
}

func ActorSchedulePost(c *gin.Context) {
	actorID, _ := strconv.ParseUint(c.PostForm("actor_id"), 10, 32)
	scriptID, _ := strconv.ParseUint(c.PostForm("script_id"), 10, 32)
	scheduleDate, err := time.Parse("2006-01-02", c.PostForm("schedule_date"))
	if err != nil {
		scheduleDate = time.Now()
	}
	schedule := &models.ActorSchedule{
		ActorID:      uint(actorID),
		ScriptID:     uint(scriptID),
		RoleName:     c.PostForm("role_name"),
		ScheduleDate: scheduleDate,
		Status:       c.PostForm("status"),
	}
	if schedule.Status == "" {
		schedule.Status = "scheduled"
	}
	if err := services.CreateActorSchedule(schedule); err != nil {
		c.Redirect(http.StatusFound, "/actors/schedule")
		return
	}
	c.Redirect(http.StatusFound, "/actors/schedule")
}
