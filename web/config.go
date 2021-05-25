package web

import (
	config2 "github.com/ovargasmahisoft/kmn-commons/config"
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
		applicationConfig := config2.Config()
		if err := applicationConfig.UnmarshalKey("server", c); err != nil {
			panic(err)
		}
		server = *c
	})
	return server
}
