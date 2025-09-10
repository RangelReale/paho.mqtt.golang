package mqtt

import "context"

type withContextData[T any] struct {
	ctx context.Context
	v   T
}

func withContext[T any](ctx context.Context, value T) withContextData[T] {
	return withContextData[T]{
		ctx: ctx,
		v:   value,
	}
}

func withoutContext[T any](value T) withContextData[T] {
	return withContextData[T]{
		ctx: context.Background(),
		v:   value,
	}
}
