package main

import (
	"github.com/gin-gonic/gin"
	bus2 "github.com/ovargasmahisoft/kmn-commons/bus"
	"github.com/ovargasmahisoft/kmn-commons/examples/internal/controllers"
	"github.com/ovargasmahisoft/kmn-commons/examples/internal/events"
	"github.com/ovargasmahisoft/kmn-commons/examples/internal/services"
	mysql2 "github.com/ovargasmahisoft/kmn-commons/migration/mysql"
	web2 "github.com/ovargasmahisoft/kmn-commons/web"
	authorization2 "github.com/ovargasmahisoft/kmn-commons/web/authorization"
	jwt2 "github.com/ovargasmahisoft/kmn-commons/web/authorization/jwt"
	"reflect"
)

func main() {

	bus := bus2.DefaultBus()
	bus.Subscribe(reflect.TypeOf(events.DummyCreatedEvent{}), events.OnDummyCreated)

	mysql2.MigrateMySql("default")
	host := web2.NewDefault().
		Use(gin.Logger()).
		UseDefaultErrorHandler().
		Use(jwt2.New().AuthenticationHandler)

	host.
		RegisterPath("/api").
		Use(authorization2.RequireAuthenticationHandler).
		RegisterController(
			controllers.NewDummyController(services.NewDummyService(bus)),
		)

	host.Start()
}
