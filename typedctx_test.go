package typedctx

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

type keyType1[T any] string

type keyType2[T any] string

func (k keyType1[T]) ContextKey(T) {}

func (k keyType2[T]) ContextKey(T) {}

const (
	stringValueCtxKey1 keyType1[string] = "string"
	intValueCtxKey1    keyType1[int]    = "int"

	stringValueCtxKey2 keyType2[string] = "string"
	intValueCtxKey2    keyType2[int]    = "int"
)

func TestTypedKeyEqual(t *testing.T) {
	baseCtx := context.Background()

	ctx := Set[string](baseCtx, stringValueCtxKey1, "str1")
	ctx = Set[int](ctx, intValueCtxKey1, 1)

	strV1, ok := Get[string](ctx, stringValueCtxKey1)
	assert.True(t, ok)
	assert.Equal(t, "str1", strV1)
	strV2, ok := Get[string](ctx, stringValueCtxKey2)
	assert.False(t, ok)
	assert.Equal(t, "", strV2)
	intV1, ok := Get[int](ctx, intValueCtxKey1)
	assert.True(t, ok)
	assert.Equal(t, 1, intV1)
	intV2, ok := Get[int](ctx, intValueCtxKey2)
	assert.False(t, ok)
	assert.Equal(t, 0, intV2)

	ctx = Set[string](ctx, stringValueCtxKey2, "str2")
	ctx = Set[int](ctx, intValueCtxKey2, 2)

	strV1, ok = Get[string](ctx, stringValueCtxKey1)
	assert.True(t, ok)
	assert.Equal(t, "str1", strV1)
	strV2, ok = Get[string](ctx, stringValueCtxKey2)
	assert.True(t, ok)
	assert.Equal(t, "str2", strV2)
	intV1, ok = Get[int](ctx, intValueCtxKey1)
	assert.True(t, ok)
	assert.Equal(t, 1, intV1)
	intV2, ok = Get[int](ctx, intValueCtxKey2)
	assert.True(t, ok)
	assert.Equal(t, 2, intV2)
}
