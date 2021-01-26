package services

import (
	"fmt"
	"github.com/ovargasmahisoft/kmn-commons/examples/internal/dtos"
	"github.com/ovargasmahisoft/kmn-commons/examples/internal/events"
	"github.com/ovargasmahisoft/kmn-commons/pkg/web"
	"github.com/ovargasmahisoft/kmn-commons/pkg/web/errors"
)

type dummyService struct {
	publisher Publisher
}

func NewDummyService(publisher Publisher) *dummyService {
	return &dummyService{
		publisher: publisher,
	}
}

func (d dummyService) Get(ctx *web.Context, id int64) (*dtos.Dummy, error) {
	return &dtos.Dummy{
		ID:    id,
		Title: fmt.Sprintf("Hello %s", *ctx.Principal().Identity().Name()),
	}, nil
}

func (d dummyService) Create(ctx *web.Context, request dtos.DummyCreateRequest) (*dtos.Dummy, error) {

	dummy := dtos.Dummy{
		ID:        1,
		Title:     request.Title,
		CompanyID: 3,
	}
	d.publisher.PublishAsync(ctx, events.DummyCreatedEvent{Dummy: dummy})
	return &dummy, nil
}

func (d dummyService) Patch(ctx *web.Context, id int64, request dtos.DummyPatchRequest) (*dtos.Dummy, error) {
	return nil, errors.NotImplementedError{
		Method:    "Patch",
		Interface: "dummyService",
	}
}

func (d dummyService) Delete(ctx *web.Context, id int64) error {
	return errors.NotImplementedError{
		Interface: "dummyService",
		Method:    "Delete",
	}
}
