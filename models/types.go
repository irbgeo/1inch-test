package models

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
)

type In struct {
	PoolID   common.Address
	TokenIn  common.Address
	TokenOut common.Address
	AmountIn decimal.Decimal
}

type Out struct {
	AmountOut decimal.Decimal
}
