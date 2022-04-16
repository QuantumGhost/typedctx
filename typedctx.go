package typedctx

import (
	"context"
)

type ContextKey[T any] interface {
	ContextKey(T)
}

func Get[T any, K ContextKey[T]](ctx context.Context, key K) (T, bool) {
	var value T
	v := ctx.Value(key)
	if v == nil {
		return value, false
	}

	t, ok := v.(T)
	if !ok {
		return value, false
	}
	value = t
	return value, true
}

func Set[T any](ctx context.Context, key ContextKey[T], value T) context.Context {
	return context.WithValue(ctx, key, value)
}
