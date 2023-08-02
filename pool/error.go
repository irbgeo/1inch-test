package pool

import "errors"

var (
	errInsufficientInputAmount = errors.New("INSUFFICIENT_INPUT_AMOUNT")
	errInsufficientLiquidity   = errors.New("INSUFFICIENT_LIQUIDITY")

	errGetPoolInfo = errors.New("get pool info")
)
