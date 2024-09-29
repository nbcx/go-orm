package opentelemetry

import (
	"bytes"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/trace"
	otelTrace "go.opentelemetry.io/otel/trace"

	"github.com/nbcx/go-orm/orm"
)

func TestFilterChainBuilderFilterChain(t *testing.T) {
	// Init Trace
	buf := bytes.NewBuffer([]byte{})
	exp, err := stdouttrace.New(stdouttrace.WithWriter(buf))
	if err != nil {
		t.Error(err)
	}
	tp := trace.NewTracerProvider(trace.WithBatcher(exp))
	otel.SetTracerProvider(tp)

	// Build FilterChain
	csf := func(ctx context.Context, span otelTrace.Span, inv *orm.Invocation) {
		span.SetAttributes(attribute.String("hello", "work"))
	}
	builder := NewFilterChainBuilder(WithCustomSpanFunc(csf))

	inv := &orm.Invocation{Method: "Hello"}
	next := func(ctx context.Context, inv *orm.Invocation) []interface{} { return nil }

	builder.FilterChain(next)(context.Background(), inv)

	// Close tp
	err = tp.Shutdown(context.Background())
	if err != nil {
		t.Error(err)
	}

	// Assert opentelemetry span name
	assert.Equal(t, "Hello#", string(buf.Bytes()[9:15]))
}
