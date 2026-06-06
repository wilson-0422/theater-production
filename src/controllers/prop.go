package controllers

import (
	"net/http"
	"strconv"
	"time"

	"theater-production/src/models"
	"theater-production/src/services"

	"github.com/gin-gonic/gin"
)

func PropList(c *gin.Context) {
	props, err := services.GetAllProps()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "props/list.html", gin.H{
			"Title": "道具服装",
			"Error": "获取道具列表失败",
		})
		return
	}
	user := GetSessionUser(c)
	c.HTML(http.StatusOK, "props/list.html", gin.H{
		"Title": "道具服装",
		"Props": props,
		"User":  user,
	})
}

func PropDetail(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.Redirect(http.StatusFound, "/props/")
		return
	}
	prop, err := services.GetPropByID(uint(id))
	if err != nil {
		c.Redirect(http.StatusFound, "/props/")
		return
	}
	requisitions, _ := services.GetPropRequisitionsByPropID(uint(id))
	user := GetSessionUser(c)
	c.HTML(http.StatusOK, "props/detail.html", gin.H{
		"Title":         prop.Name,
		"Prop":          prop,
		"Requisitions":  requisitions,
		"User":          user,
	})
}

func PropRequisitionView(c *gin.Context) {
	props, _ := services.GetAllProps()
	actors, _ := services.GetAllActors()
	requisitions, _ := services.GetAllPropRequisitions()
	user := GetSessionUser(c)
	c.HTML(http.StatusOK, "props/requisition.html", gin.H{
		"Title":         "道具服装领用",
		"Props":         props,
		"Actors":        actors,
		"Requisitions":  requisitions,
		"User":          user,
	})
}

func PropRequisitionPost(c *gin.Context) {
	propID, _ := strconv.ParseUint(c.PostForm("prop_id"), 10, 32)
	actorID, _ := strconv.ParseUint(c.PostForm("actor_id"), 10, 32)
	quantity, _ := strconv.Atoi(c.PostForm("quantity"))
	reqDate, err := time.Parse("2006-01-02", c.PostForm("requisition_date"))
	if err != nil {
		reqDate = time.Now()
	}
	var returnDate time.Time
	returnStr := c.PostForm("return_date")
	if returnStr != "" {
		returnDate, _ = time.Parse("2006-01-02", returnStr)
	}
	requisition := &models.PropRequisition{
		PropID:          uint(propID),
		ActorID:         uint(actorID),
		Quantity:        quantity,
		RequisitionDate: reqDate,
		ReturnDate:      returnDate,
		Status:          "pending",
		Notes:           c.PostForm("notes"),
	}
	if err := services.CreatePropRequisition(requisition); err != nil {
		c.Redirect(http.StatusFound, "/props/requisition")
		return
	}
	c.Redirect(http.StatusFound, "/props/requisition")
}
