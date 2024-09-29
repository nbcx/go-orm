package opentracing

import (
	"context"
	"testing"
	"time"

	"github.com/opentracing/opentracing-go"

	"github.com/nbcx/go-orm/orm"
)

func TestFilterChainBuilderFilterChain(t *testing.T) {
	next := func(ctx context.Context, inv *orm.Invocation) []interface{} {
		inv.TxName = "Hello"
		return []interface{}{}
	}

	builder := &FilterChainBuilder{
		CustomSpanFunc: func(span opentracing.Span, ctx context.Context, inv *orm.Invocation) {
			span.SetTag("hello", "hell")
		},
	}

	inv := &orm.Invocation{
		Method:      "Hello",
		TxStartTime: time.Now(),
	}
	builder.FilterChain(next)(context.Background(), inv)
}
