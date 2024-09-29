package opentelemetry

import (
	"context"
	"strings"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	otelTrace "go.opentelemetry.io/otel/trace"

	"github.com/nbcx/go-orm/orm"
)

type (
	CustomSpanFunc    func(ctx context.Context, span otelTrace.Span, inv *orm.Invocation)
	FilterChainOption func(fcv *FilterChainBuilder)
)

// FilterChainBuilder provides an opentelemtry Filter
type FilterChainBuilder struct {
	// customSpanFunc users are able to custom their span
	customSpanFunc CustomSpanFunc
}

func NewFilterChainBuilder(options ...FilterChainOption) *FilterChainBuilder {
	fcb := &FilterChainBuilder{}
	for _, o := range options {
		o(fcb)
	}
	return fcb
}

// WithCustomSpanFunc add function to custom span
func WithCustomSpanFunc(customSpanFunc CustomSpanFunc) FilterChainOption {
	return func(fcv *FilterChainBuilder) {
		fcv.customSpanFunc = customSpanFunc
	}
}

// FilterChain traces invocation with opentelemetry
// Unless invocation.Method is Begin*, Commit or Rollback
func (builder *FilterChainBuilder) FilterChain(next orm.Filter) orm.Filter {
	return func(ctx context.Context, inv *orm.Invocation) []interface{} {
		if strings.HasPrefix(inv.Method, "Begin") || inv.Method == "Commit" || inv.Method == "Rollback" {
			return next(ctx, inv)
		}
		spanCtx, span := otel.Tracer("beego_orm").Start(ctx, invOperationName(ctx, inv))
		defer span.End()

		res := next(spanCtx, inv)
		builder.buildSpan(spanCtx, span, inv)
		return res
	}
}

// buildSpan add default span attributes and custom attributes with customSpanFunc
func (builder *FilterChainBuilder) buildSpan(ctx context.Context, span otelTrace.Span, inv *orm.Invocation) {
	span.SetAttributes(attribute.String("orm.method", inv.Method))
	span.SetAttributes(attribute.String("orm.table", inv.GetTableName()))
	span.SetAttributes(attribute.Bool("orm.insideTx", inv.InsideTx))
	v, _ := ctx.Value(orm.TxNameKey).(string)
	span.SetAttributes(attribute.String("orm.txName", v))
	span.SetAttributes(attribute.String("span.kind", "client"))
	span.SetAttributes(attribute.String("component", "beego"))

	if builder.customSpanFunc != nil {
		builder.customSpanFunc(ctx, span, inv)
	}
}

func invOperationName(ctx context.Context, inv *orm.Invocation) string {
	if n, ok := ctx.Value(orm.TxNameKey).(string); ok {
		return inv.Method + "#tx(" + n + ")"
	}
	return inv.Method + "#" + inv.GetTableName()
}
