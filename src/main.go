package main

import (
	"theater-production/src/config"
	"theater-production/src/routes"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"html/template"
	"time"
)

func formatTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func formatDate(t time.Time) string {
	return t.Format("2006-01-02")
}

func main() {
	config.InitDB()

	Run()

	r := gin.Default()

	funcMap := template.FuncMap{
		"formatTime":  formatTime,
		"formatDate":  formatDate,
		"safeHTML":    func(s string) template.HTML { return template.HTML(s) },
	}
	r.SetFuncMap(funcMap)
	r.LoadHTMLGlob("templates/**/*")
	r.Static("/static", "./static")

	store := cookie.NewStore([]byte(config.AppConfig.SessionSecret))
	r.Use(sessions.Sessions("theater_session", store))

	routes.SetupRoutes(r)

	r.Run(":" + config.AppConfig.Port)
}
