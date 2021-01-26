package async

import "context"

type Task interface {
	Await() (interface{}, error)
}

type task struct {
	await func(ctx context.Context) (interface{}, error)
}

func (f task) Await() (interface{}, error) {
	return f.await(context.Background())
}

func ExecTask(f func() (interface{}, error)) Task {
	var result interface{}
	var err error
	c := make(chan struct{})
	go func() {
		defer close(c)
		result, err = f()
	}()
	return task{
		await: func(ctx context.Context) (interface{}, error) {
			select {
			case <-ctx.Done():
				return nil, ctx.Err()
			case <-c:
				return result, err
			}
		},
	}
}
