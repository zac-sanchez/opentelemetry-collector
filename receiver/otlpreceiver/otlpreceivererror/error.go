package otlpreceivererror

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type BackpressureError struct{}

func (b BackpressureError) Error() string {
	return "backpressure error"
}

func (b BackpressureError) GRPCStatus() *status.Status {
	return status.New(codes.ResourceExhausted, "backpressure error")
}
