package models

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
)

type IPool interface {
	GetToken0() common.Address
	GetToken1() common.Address
	GetAmountOut(fromToken common.Address, inputAmount decimal.Decimal) (decimal.Decimal, error)
}
