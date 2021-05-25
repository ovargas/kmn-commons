package services

import (
	"context"
	async2 "github.com/ovargasmahisoft/kmn-commons/async"
	bus2 "github.com/ovargasmahisoft/kmn-commons/bus"
)

type Publisher interface {
	PublishAsync(ctx context.Context, event bus2.Event) async2.Action
}
