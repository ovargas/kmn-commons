package bus

import (
	"context"
	"github.com/ovargasmahisoft/kmn-commons/pkg/async"
	"reflect"
	"strings"
	"sync"
)

type Event interface {
}

type publishError struct {
	Errors []error
}

func (e publishError) Error() string {
	if e.Errors == nil || len(e.Errors) == 0 {
		return ""
	}
	var errors []string
	for _, err := range e.Errors {
		errors = append(errors, err.Error())
	}
	return strings.Join(errors, ",")
}

type Handler func(ctx context.Context, event Event) error
type handlers []Handler
type ApplicationBus map[reflect.Type]handlers

func (l ApplicationBus) Subscribe(eventType reflect.Type, handler ...Handler) {
	if _, ok := l[eventType]; !ok {
		l[eventType] = handlers{}
	}

	for _, h := range handler {
		pointer := reflect.ValueOf(h).Pointer()
		exist := false
		for _, i := range l[eventType] {
			exist = pointer == reflect.ValueOf(i).Pointer()
		}
		if !exist {
			l[eventType] = append(l[eventType], h)
		}
	}
}

func (l ApplicationBus) PublishAsync(ctx context.Context, event Event) async.Action {
	var futures []async.Action
	if value, ok := l[reflect.TypeOf(event)]; ok {
		for _, h := range value {
			h := h
			future := async.ExecAction(func() error {
				return h(ctx, event)
			})
			futures = append(futures, future)
		}
	}

	return async.ExecAction(func() error {

		var errors []error

		for _, f := range futures {
			err := f.Await()
			if err != nil {
				errors = append(errors, err)
			}
		}

		if errors != nil {
			return publishError{
				Errors: errors,
			}
		}

		return nil
	})
}

var (
	eventHandlers *ApplicationBus
	once          sync.Once
)

func DefaultBus() *ApplicationBus {
	once.Do(func() {
		eventHandlers = &ApplicationBus{}
	})
	return eventHandlers
}
