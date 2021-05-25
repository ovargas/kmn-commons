package web

import (
	"context"
	authorization2 "github.com/ovargasmahisoft/kmn-commons/web/authorization"
)

type Context struct {
	context.Context
	principal authorization2.IPrincipal
}

func (c Context) Principal() authorization2.IPrincipal {
	return c.principal
}

func ApplicationContext(principal *authorization2.IPrincipal) Context {

	var p authorization2.IPrincipal
	if principal == nil {
		p = authorization2.Anonymous()
	} else {
		p = *principal
	}
	return Context{
		Context:   context.Background(),
		principal: p,
	}
}
