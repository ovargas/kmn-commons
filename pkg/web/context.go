package web

import (
	"context"
	"github.com/ovargasmahisoft/kmn-commons/pkg/web/authorization"
)

type Context struct {
	context.Context
	principal authorization.IPrincipal
}

func (c Context) Principal() authorization.IPrincipal {
	return c.principal
}

func ApplicationContext(principal *authorization.IPrincipal) Context {

	var p authorization.IPrincipal
	if principal == nil {
		p = authorization.Anonymous()
	} else {
		p = *principal
	}
	return Context{
		Context:   context.Background(),
		principal: p,
	}
}
