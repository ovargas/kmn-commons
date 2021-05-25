package bus

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

type Dummy struct {
}

func TestBus(t *testing.T) {
	bus := DefaultBus()
	handler1Called := false
	handler2Called := false

	handler1 := func(ctx context.Context, event Event) error {
		handler1Called = true
		return nil
	}

	handler2 := func(ctx context.Context, event Event) error {
		handler2Called = true
		return nil
	}

	handlerError1 := func(ctx context.Context, event Event) error {
		return fmt.Errorf("dummy error")
	}

	bus.Subscribe(reflect.TypeOf(Dummy{}), handler1, handler2, handlerError1)

	err := bus.PublishAsync(context.Background(), Dummy{}).Await()

	assert.True(t, handler1Called, "Handler 1 should have been called")
	assert.True(t, handler2Called, "Handler 2 should have been called")
	assert.Equal(t, err.Error(), "dummy error")
}
