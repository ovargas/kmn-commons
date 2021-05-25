package async

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAsyncActionExecution(t *testing.T) {

	a := ExecAction(func() error {
		return nil
	})

	err := a.Await()

	assert.Nil(t, err, "error should be nil")
}

func TestAsyncActionExecutionReturningError(t *testing.T) {

	a := ExecAction(func() error {
		return fmt.Errorf("dummy error")
	})

	err := a.Await()

	assert.NotNil(t, err, "error shouldn't be nil")
	assert.Equal(t, err.Error(), "dummy error")
}


func TestAsyncTaskExecution(t *testing.T) {
	a := ExecTask(func() (interface{}, error) {
		return struct {

		}{}, nil
	})

	result, err := a.Await()

	assert.NotNil(t, result, "result shouldn't be nil")
	assert.Nil(t, err, "error should be nil")
}

func TestAsyncTaskExecutionReturningError(t *testing.T) {
	a := ExecTask(func() (interface{}, error) {
		return nil, fmt.Errorf("dummy error")
	})

	result, err := a.Await()

	assert.NotNil(t, err, "error shouldn't be nil")
	assert.Equal(t, err.Error(), "dummy error")
	assert.Nil(t, result, "result should be nil")
}
