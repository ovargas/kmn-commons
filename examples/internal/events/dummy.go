package events

import (
	"context"
	"fmt"
	"github.com/ovargasmahisoft/kmn-commons/pkg/bus"
)

func OnDummyCreated(ctx context.Context, event bus.Event) error {
	fmt.Printf("Dummy %v created\n", event)
	return nil
}
