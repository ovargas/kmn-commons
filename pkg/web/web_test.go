package web

import (
	"github.com/gin-gonic/gin"
	"github.com/ovargasmahisoft/kmn-commons/pkg/config"
	"github.com/ovargasmahisoft/kmn-commons/pkg/web/authorization"
	"github.com/ovargasmahisoft/kmn-commons/pkg/web/authorization/jwt"
	"github.com/stretchr/testify/assert"
	"os"
	"path"
	"runtime"
	"testing"
)

func init() {
	_, filename, _, _ := runtime.Caller(0)
	os.Setenv(config.EnvConfigPath, path.Join(path.Dir(filename), "../../")+"/test-resources/")
}

type DummyController struct {
}

func (d DummyController) Register(router *gin.RouterGroup) {
	r := router.Group("/dummy")
	r.GET("", func(context *gin.Context) {

	})
	r.POST("", func(context *gin.Context) {

	})
}

func TestWebHostBuilder(t *testing.T) {

	host := NewDefault().
		Use(gin.Logger()).
		Use(jwt.New().AuthenticationHandler).
		UseDefaultErrorHandler()

	host.
		RegisterPath("/engine").
		Use(authorization.RequireAuthenticationHandler).
		RegisterController(
			&DummyController{},
		)

	assert.Equal(t, 3, len(host.engine.Handlers))
	assert.Equal(t, 2, len(host.engine.Routes()))
	assert.Equal(t, 3000, host.port)
}
