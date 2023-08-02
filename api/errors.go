package api

import "errors"

var (
	errPoolIDNumber      = errors.New("the number of poolIDs should be 1")
	errFromTokenNumber   = errors.New("the number of fromTokens should be 1")
	errInputAmountNumber = errors.New("the number of inputAmounts should be 1")
	errParseInputAmount  = errors.New("failed to parse amount in")
	errToTokenNumber     = errors.New("the number of toTokens should be 1")
)
