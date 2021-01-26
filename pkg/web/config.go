package web

import (
	"github.com/ovargasmahisoft/kmn-commons/pkg/config"
	"sync"
)

type Server struct {
	Port     int
	HostName string
}

var (
	onceConfiguration sync.Once
	server            Server
)

func Config() Server {
	onceConfiguration.Do(func() {
		c := &Server{}
		applicationConfig := config.Config()
		if err := applicationConfig.UnmarshalKey("server", c); err != nil {
			panic(err)
		}
		server = *c
	})
	return server
}
