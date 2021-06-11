package time_series

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var OutOfRange = status.Error(codes.OutOfRange, "out of range")
var InvalidArgument = status.Error(codes.InvalidArgument, "invalid argument")
