package authorization

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//RequireAuthenticationHandler middleware to ensure the user is authenticated
func RequireAuthenticationHandler(ctx *gin.Context) {
	if principal, ok := ctx.Get(PrincipalContextName); !ok || principal.(IPrincipal).Identity().AuthenticationType() == "ANONYMOUS" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}
	ctx.Next()
}
