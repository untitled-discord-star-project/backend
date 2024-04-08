package ctxutil

import "context"

type key[T any] struct{}

func WithValue[T any](ctx context.Context, value T) context.Context {
	return context.WithValue(ctx, key[T]{}, value)
}

func Value[T any](ctx context.Context) (T, bool) {
	value, ok := ctx.Value(key[T]{}).(T)
	return value, ok
}
