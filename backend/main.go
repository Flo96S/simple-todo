package main

import (
	"dsystem/backend/pck/database"
	"dsystem/backend/pck/helper"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/trace"
)

var tracer trace.Tracer

const meterName = ""

func main() {
	fmt.Println("Hello")
	database.Init()
	r := gin.Default()

	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, helper.GenerateRandomId(32))
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Working")
	})

	r.GET("/posts", func(c *gin.Context) {
		list := database.GetAllItemsShort()
		c.JSON(http.StatusOK, list)
	})

	r.GET("/post/:id", func(c *gin.Context) {
		item := database.GetItem(c.Param("id"))
		c.JSON(http.StatusOK, item)
	})

	r.POST("/post", func(c *gin.Context) {
		success := database.InsertItem(c)
		if success {
			c.JSON(http.StatusOK, "Done")
			return
		}
		c.JSON(http.StatusBadRequest, "Not successful")
	})

	r.PUT("/post", func(c *gin.Context) {
		id := c.PostForm("id")
		success := database.Update(id, true)
		if success {
			c.JSON(http.StatusOK, "")
			return
		}
		c.JSON(http.StatusBadRequest, "")
	})

	r.DELETE("/post/:id", func(c *gin.Context) {
		success := database.DeleteItem(c)
		if success {
			c.JSON(http.StatusOK, "Deleted")
			return
		}
		c.JSON(http.StatusBadRequest, "Still exists")
	})

	r.Run(":9988")
}
