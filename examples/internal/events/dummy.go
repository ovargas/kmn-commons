package events

import (
	"context"
	"fmt"
	bus2 "github.com/ovargasmahisoft/kmn-commons/bus"
)

func OnDummyCreated(ctx context.Context, event bus2.Event) error {
	fmt.Printf("Dummy %v created\n", event)
	return nil
}
