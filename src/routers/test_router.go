package routers

import (
	"github.com/gin-gonic/gin"
	"golan-clean-web-api/api/handlers"
)

func TestRouter(r *gin.RouterGroup) {

	h := handlers.NewTestHandler()
	r.GET("/", h.Test)
	r.GET("/user/:id", h.UserById)
	r.GET("/binder/h1", h.HeaderBinder1)
	r.GET("/binder/h2", h.HeaderBinder2)
	r.GET("/binder/query", h.QueryBinder1)
	r.GET("/binder/body", h.BodyBinder)
}
