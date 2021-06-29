package errors

import "errors"

var InvalidAmount = errors.New("invalid amount")
var InsufficientFunds = errors.New("insufficient funds")
var InvalidArgument = errors.New("invalid argument")
