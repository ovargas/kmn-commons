package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ovargasmahisoft/kmn-commons/pkg/web/errors"
	"net/http"
)

//ErrorHandler check for errors and returns true if
//the error was handled
type ErrorHandler func(ctx *gin.Context) bool

type Controller interface {
	Register(router *gin.RouterGroup)
}

type IHostPath interface {
	Use(middleware ...gin.HandlerFunc) IHostPath
	RegisterController(controller ...Controller) IHostPath
}

type Host struct {
	engine   *gin.Engine
	hostName string
	port     int
}

func New(server Server) Host {
	return Host{
		engine:   gin.New(),
		hostName: server.HostName,
		port:     server.Port,
	}
}

func (h Host) Engine() *gin.Engine {
	return h.engine
}

func NewDefault() Host {
	return New(Config())
}

func (h Host) Use(middleware ...gin.HandlerFunc) Host {
	h.engine.Use(middleware...)
	return h
}

func (h Host) UseDefaultErrorHandler() Host {
	h.ErrorHandlers(errors.DefaultErrorHandler)
	return h
}

func (h Host) ErrorHandlers(errorHandler ...ErrorHandler) {
	h.engine.Use(func(ctx *gin.Context) {
		ctx.Next()
		if len(ctx.Errors) == 0 {
			return
		}

		for _, handler := range errorHandler {
			if handler(ctx) {
				return
			}
		}

		ctx.AbortWithStatusJSON(http.StatusInternalServerError, &errors.ApiError{
			Title: http.StatusText(http.StatusInternalServerError),
		})
	})
}

func (h Host) RegisterPath(path string) IHostPath {
	return newHostPath(h.engine.Group(path))
}

func (h Host) Start() {
	err := h.engine.Run(fmt.Sprintf("%s:%d", h.hostName, h.port))
	panic(err)
}
