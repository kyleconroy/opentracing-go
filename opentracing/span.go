package opentracing

import "golang.org/x/net/context"

type Span interface {
	// Creates and starts a child span.
	//
	// For more information about `keyValueTags`, see the documentation for
	// `OpenTracer.StartTrace()`.
	StartChild(operationName string, keyValueTags ...interface{}) Span

	// Adds a tag to the span. The `value` is immediately coerced into a string
	// using fmt.Sprint().
	//
	// If there is a pre-existing tag set for `key`, it is overwritten.
	SetTag(key string, value interface{}) Span

	// `Message` is a format string and can refer to fields in the payload by path, like so:
	//
	//   "first transaction is worth ${transactions[0].amount} ${transactions[0].currency}"
	//
	// , and the payload might look something like
	//
	//   map[string]interface{}{
	//       transactions: map[string]interface{}[
	//           {amount: 10, currency: "USD"},
	//           {amount: 11, currency: "USD"},
	//       ]}
	Info(message string, payload ...interface{})

	// Like Info(), but for errors.
	Error(message string, payload ...interface{})

	// Sets the end timestamp and calls the `Recorder`s RecordSpan()
	// internally.
	//
	// Finish() should be the last call made to any span instance, and to do
	// otherwise leads to undefined behavior.
	Finish()

	// Suitable for serializing over the wire, etc.
	TraceContext() TraceContext

	// A convenience method. Equivalent to
	//
	//    var goCtx context.Context = ...
	//    var span Span = ...
	//    goCtx := opentracing.GoContextWithSpan(ctx, span)
	//
	AddToGoContext(goCtx context.Context) (Span, context.Context)
}