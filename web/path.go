package web

import "github.com/gin-gonic/gin"

type hostPath struct {
	router *gin.RouterGroup
}

func newHostPath(router *gin.RouterGroup) IHostPath {
	return hostPath{
		router: router,
	}
}

func (h hostPath) Use(middleware ...gin.HandlerFunc) IHostPath {
	h.router.Use(middleware...)
	return h
}

func (h hostPath) RegisterController(controller ...Controller) IHostPath {
	for _, c := range controller {
		c.Register(h.router)
	}
	return h
}
