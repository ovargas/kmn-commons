package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ovargasmahisoft/kmn-commons/examples/internal/controllers"
	"github.com/ovargasmahisoft/kmn-commons/examples/internal/events"
	"github.com/ovargasmahisoft/kmn-commons/examples/internal/services"
	"github.com/ovargasmahisoft/kmn-commons/pkg/bus"
	"github.com/ovargasmahisoft/kmn-commons/pkg/migration/mysql"
	"github.com/ovargasmahisoft/kmn-commons/pkg/web"
	"github.com/ovargasmahisoft/kmn-commons/pkg/web/authorization"
	"github.com/ovargasmahisoft/kmn-commons/pkg/web/authorization/jwt"
	"reflect"
)

func main() {

	bus := bus.DefaultBus()
	bus.Subscribe(reflect.TypeOf(events.DummyCreatedEvent{}), events.OnDummyCreated)

	mysql.MigrateMySql("default")
	host := web.NewDefault().
		Use(gin.Logger()).
		UseDefaultErrorHandler().
		Use(jwt.New().AuthenticationHandler)

	host.
		RegisterPath("/api").
		Use(authorization.RequireAuthenticationHandler).
		RegisterController(
			controllers.NewDummyController(services.NewDummyService(bus)),
		)

	host.Start()
}
