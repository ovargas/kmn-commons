package services

import (
	"context"
	"github.com/ovargasmahisoft/kmn-commons/pkg/async"
	"github.com/ovargasmahisoft/kmn-commons/pkg/bus"
)

type Publisher interface {
	PublishAsync(ctx context.Context, event bus.Event) async.Action
}
