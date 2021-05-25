package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ovargasmahisoft/kmn-commons/examples/internal/dtos"
	web2 "github.com/ovargasmahisoft/kmn-commons/web"
	"net/http"
	"strconv"
)

type DummyService interface {
	Get(ctx *web2.Context, id int64) (*dtos.Dummy, error)
	Create(ctx *web2.Context, request dtos.DummyCreateRequest) (*dtos.Dummy, error)
	Patch(ctx *web2.Context, id int64, request dtos.DummyPatchRequest) (*dtos.Dummy, error)
	Delete(ctx *web2.Context, id int64) error
}

type DummyController struct {
	service DummyService
}

func NewDummyController(service DummyService) *DummyController {
	return &DummyController{
		service: service,
	}
}

func (dc DummyController) Register(router *gin.RouterGroup) {
	r := router.Group("dummies")
	r.GET("/:dummyID", dc.get)
	r.POST("", dc.create)
	r.PATCH("/:dummyID", dc.patch)
	r.DELETE("/:dummyID", dc.delete)
}

func (dc DummyController) get(ctx *gin.Context) {
	dummyID, err := strconv.ParseInt(ctx.Param("dummyID"), 10, 64)

	if err != nil {
		ctx.Error(err)
		return
	}

	dummy, err := dc.service.Get(newApplicationContext(ctx), dummyID)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(200, dummy)
}

func (dc DummyController) create(ctx *gin.Context) {

	var request = dtos.DummyCreateRequest{}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.Error(err)
		return
	}

	if response, err := dc.service.Create(newApplicationContext(ctx), request); err == nil {
		ctx.JSON(http.StatusCreated, response)
	} else {
		ctx.Error(err)
	}

}

func (dc DummyController) patch(ctx *gin.Context) {
	dummyID, err := strconv.ParseInt(ctx.Param("dummyID"), 10, 64)
	if err != nil {
		ctx.Error(err)
		return
	}

	var request = dtos.DummyPatchRequest{}

	if err := ctx.ShouldBind(&request); err != nil {
		ctx.Error(err)
		return
	}

	if response, err := dc.service.Patch(newApplicationContext(ctx), dummyID, request); err == nil {
		ctx.JSON(http.StatusOK, response)
	} else {
		ctx.Error(err)
	}

}

func (dc DummyController) delete(ctx *gin.Context) {
	dummyID, err := strconv.ParseInt(ctx.Param("dummyID"), 10, 64)
	if err != nil {
		ctx.Error(err)
		return
	}

	if err := dc.service.Delete(newApplicationContext(ctx), dummyID); err == nil {
		ctx.JSON(http.StatusNoContent, http.NoBody)
	} else {
		ctx.Error(err)
	}
}
