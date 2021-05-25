package async

import "context"

type Action interface {
	Await() error
}

type action struct {
	await func(ctx context.Context) error
}

func (f action) Await() error {
	return f.await(context.Background())
}

func ExecAction(f func() error) Action {
	var err error
	c := make(chan struct{})
	go func() {
		defer close(c)
		err = f()
	}()
	return action{
		await: func(ctx context.Context) error {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-c:
				return err
			}
		},
	}
}
