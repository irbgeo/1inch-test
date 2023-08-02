package models

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
)

type In struct {
	PoolID      common.Address
	FromToken   common.Address
	ToToken     common.Address
	InputAmount decimal.Decimal
}

type Out struct {
	AmountOut decimal.Decimal
}
