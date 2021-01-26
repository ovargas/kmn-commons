package events

import "github.com/ovargasmahisoft/kmn-commons/examples/internal/dtos"

type DummyCreatedEvent struct {
	dtos.Dummy
}

type DummyPatchedEvent struct {
	dtos.Dummy
}