package main

import (
	"backend/pck/database"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/trace"
)

var tracer trace.Tracer

const meterName = ""

func main() {
	// Set up OpenTelemetry.
	/*
		serviceName := "Todo"
		serviceVersion := "0.1"
		otelShutdown, err := otel.SetupOTelSDK(context.Background(), serviceName, serviceVersion)
		if err != nil {
			return
		}
		defer func() {
			err = errors.Join(err, otelShutdown(context.Background()))
		}()
	*/
	database.Init()
	r := gin.Default()
	//r.Use(gin.WrapH(otelhttp.NewHandler(http.DefaultServeMux, "/v1/trace")))

	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"success:": database.Check(),
			"error":    database.ErrorCheck(),
		})
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Working api - distributedsystems")
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
		fmt.Println("Success: ", success)
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
