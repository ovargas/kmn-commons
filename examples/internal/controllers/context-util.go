package controllers

import (
	"github.com/gin-gonic/gin"
	web2 "github.com/ovargasmahisoft/kmn-commons/web"
	authorization2 "github.com/ovargasmahisoft/kmn-commons/web/authorization"
)

func newApplicationContext(ctx *gin.Context) *web2.Context {
	c := web2.ApplicationContext(getPrincipal(ctx))
	return &c
}

func getPrincipal(ctx *gin.Context) *authorization2.IPrincipal {
	if p, ok := ctx.Get(authorization2.PrincipalContextName); ok {
		v := p.(authorization2.IPrincipal)
		return &v
	}

	return nil
}
