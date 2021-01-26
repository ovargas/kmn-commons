package errors

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestNotImplementedError(t *testing.T) {
	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Error(NotImplementedError{"Test", "errors"})
	handled := DefaultErrorHandler(ctx)
	assert.True(t, handled)
	assert.True(t, ctx.IsAborted())
	assert.Equal(t, http.StatusMethodNotAllowed, recorder.Code)
}

func TestResourceNotFoundError(t *testing.T) {
	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Error(ResourceNotFoundError{"MyId", "DummyObject"})
	handled := DefaultErrorHandler(ctx)
	assert.True(t, handled)
	assert.True(t, ctx.IsAborted())
	assert.Equal(t, http.StatusNotFound, recorder.Code)
}

func TestDuplicatedEntryError(t *testing.T) {
	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Error(DuplicatedEntryError{"MyId", "DummyObject"})
	handled := DefaultErrorHandler(ctx)
	assert.True(t, handled)
	assert.True(t, ctx.IsAborted())
	assert.Equal(t, http.StatusConflict, recorder.Code)
}

func TestForbiddenError(t *testing.T) {
	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Error(ForbiddenError{"no access"})
	handled := DefaultErrorHandler(ctx)
	assert.True(t, handled)
	assert.True(t, ctx.IsAborted())
	assert.Equal(t, http.StatusForbidden, recorder.Code)
}

func TestInvalidArgumentError(t *testing.T) {
	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Error(InvalidArgumentError{"invalid value"})
	handled := DefaultErrorHandler(ctx)
	assert.True(t, handled)
	assert.True(t, ctx.IsAborted())
	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}

func TestNumError(t *testing.T) {
	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	_, err := strconv.ParseInt("not a number", 10, 64)
	ctx.Error(err)
	handled := DefaultErrorHandler(ctx)
	assert.True(t, handled)
	assert.True(t, ctx.IsAborted())
	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}

type DummyModel struct {
	Value *string `json:"value" binding:"required" validate:"required"`
}

func TestValidationErrors(t *testing.T) {
	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	dummy := &DummyModel{}
	// Internally ctx.ShouldBindJson uses validator so it is enough
	// to test that the handler can manage the struct validation errors
	v := validator.New()
	ctx.Error(v.Struct(dummy))
	handled := DefaultErrorHandler(ctx)
	assert.True(t, handled)
	assert.True(t, ctx.IsAborted())
	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}

func TestNotHandle(t *testing.T) {
	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Error(fmt.Errorf("dummy error"))
	handled := DefaultErrorHandler(ctx)
	assert.False(t, handled)
	assert.False(t, ctx.IsAborted())
}