package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ovargasmahisoft/kmn-commons/pkg/web"
	"github.com/ovargasmahisoft/kmn-commons/pkg/web/authorization"
)

func newApplicationContext(ctx *gin.Context) *web.Context {
	c := web.ApplicationContext(getPrincipal(ctx))
	return &c
}

func getPrincipal(ctx *gin.Context) *authorization.IPrincipal {
	if p, ok := ctx.Get(authorization.PrincipalContextName); ok {
		v := p.(authorization.IPrincipal)
		return &v
	}

	return nil
}
