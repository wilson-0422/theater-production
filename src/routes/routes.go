package routes

import (
	"theater-production/src/controllers"
	"theater-production/src/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/", controllers.IndexHandler)

	auth := r.Group("/auth")
	{
		auth.GET("/login", controllers.LoginHandler)
		auth.POST("/login", controllers.LoginPost)
		auth.GET("/register", controllers.RegisterHandler)
		auth.POST("/register", controllers.RegisterPost)
	}

	authorized := r.Group("/")
	authorized.Use(middleware.AuthRequired())
	{
		authorized.GET("/dashboard", controllers.DashboardOverview)
		authorized.GET("/auth/logout", controllers.LogoutHandler)

		scripts := authorized.Group("/scripts")
		{
			scripts.GET("/", controllers.ScriptList)
			scripts.GET("/create", controllers.ScriptCreate)
			scripts.POST("/create", controllers.ScriptCreatePost)
			scripts.GET("/:id", controllers.ScriptDetail)
			scripts.GET("/:id/edit", controllers.ScriptEdit)
			scripts.POST("/:id/edit", controllers.ScriptEditPost)
			scripts.POST("/:id/delete", controllers.ScriptDelete)
		}

		actors := authorized.Group("/actors")
		{
			actors.GET("/", controllers.ActorList)
			actors.GET("/schedule", controllers.ActorScheduleView)
			actors.POST("/schedule", controllers.ActorSchedulePost)
			actors.GET("/:id", controllers.ActorDetail)
		}

		props := authorized.Group("/props")
		{
			props.GET("/", controllers.PropList)
			props.GET("/requisition", controllers.PropRequisitionView)
			props.POST("/requisition", controllers.PropRequisitionPost)
			props.GET("/:id", controllers.PropDetail)
		}

		theaters := authorized.Group("/theaters")
		{
			theaters.GET("/", controllers.TheaterList)
			theaters.GET("/schedule", controllers.TheaterScheduleView)
			theaters.POST("/schedule", controllers.TheaterSchedulePost)
			theaters.GET("/:id", controllers.TheaterDetail)
		}

		tours := authorized.Group("/tours")
		{
			tours.GET("/", controllers.TourList)
			tours.GET("/create", controllers.TourCreate)
			tours.POST("/create", controllers.TourCreatePost)
			tours.GET("/:id", controllers.TourDetail)
		}
	}
}
