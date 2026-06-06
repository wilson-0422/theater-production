package controllers

import (
	"net/http"

	"theater-production/src/services"

	"github.com/gin-gonic/gin"
)

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"Title": "剧目排演管理系统",
	})
}

func DashboardOverview(c *gin.Context) {
	scriptCount, _ := services.CountScripts()
	actorCount, _ := services.CountActors()
	propCount, _ := services.CountProps()
	theaterCount, _ := services.CountTheaters()
	tourCount, _ := services.CountTours()

	recentScripts, _ := services.GetAllScripts()
	if len(recentScripts) > 5 {
		recentScripts = recentScripts[:5]
	}

	ongoingTours, _ := services.GetToursByStatus("ongoing")
	plannedTours, _ := services.GetToursByStatus("planned")

	user := GetSessionUser(c)
	c.HTML(http.StatusOK, "dashboard/overview.html", gin.H{
		"Title":         "仪表盘",
		"ScriptCount":   scriptCount,
		"ActorCount":    actorCount,
		"PropCount":     propCount,
		"TheaterCount":  theaterCount,
		"TourCount":     tourCount,
		"RecentScripts": recentScripts,
		"OngoingTours":  ongoingTours,
		"PlannedTours":  plannedTours,
		"User":          user,
	})
}
