package mock

import (
	"context"

	"github.com/nbcx/log"
)

type mockCtxKeyType string

const mockCtxKey = mockCtxKeyType("beego-orm-mock")

func CtxWithMock(ctx context.Context, mock ...*Mock) context.Context {
	return context.WithValue(ctx, mockCtxKey, mock)
}

func mockFromCtx(ctx context.Context) []*Mock {
	ms := ctx.Value(mockCtxKey)
	if ms != nil {
		if res, ok := ms.([]*Mock); ok {
			return res
		}
		log.Error("mockCtxKey found in context, but value is not type []*Mock")
	}
	return nil
}
