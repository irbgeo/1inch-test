package api

import "errors"

var (
	errPoolIDNumber   = errors.New("the number of poolIDs should be 1")
	errTokenInNumber  = errors.New("the number of tokenIns should be 1")
	errAmountInNumber = errors.New("the number of amountIns should be 1")
	errParseAmountIn  = errors.New("failed to parse amount in")
	errTokenOutNumber = errors.New("the number of tokenOuts should be 1")
)
