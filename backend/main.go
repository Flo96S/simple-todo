package main

import (
	"backend/pck/database"
	"backend/pck/endpoint"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/trace"
)

var tracer trace.Tracer

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

	r.GET("/", endpoint.ApiRoot)
	r.GET("/test", endpoint.Test)
	r.GET("/posts", endpoint.GetTodos)
	r.GET("/post/:id", endpoint.GetTodoById)

	r.POST("/post", endpoint.AddItem)

	r.PUT("/post", endpoint.UpdateItem)

	r.DELETE("/post/:id", endpoint.DeleteItem)

	r.Run(":9988")
}
