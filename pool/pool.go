package pool

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
)

var (
	dec1k = decimal.New(1, 3)
)

type pool struct {
	token0 common.Address
	token1 common.Address

	reserve0 decimal.Decimal
	reserve1 decimal.Decimal
}

func (s *pool) GetToken0() common.Address {
	return s.token0
}

func (s *pool) GetToken1() common.Address {
	return s.token1
}

func (s *pool) GetAmountOut(fromToken common.Address, inputAmount decimal.Decimal) (decimal.Decimal, error) {
	reserveIn, reserveOut := s.getReserves(fromToken)

	return getAmountOut(inputAmount, reserveIn, reserveOut)
}

func (s *pool) getReserves(fromToken common.Address) (decimal.Decimal, decimal.Decimal) {
	if fromToken == s.token0 {
		return s.reserve0, s.reserve1
	}
	return s.reserve1, s.reserve0
}

// getAmountOut given an input amount of an asset and pair reserves
// returns the maximum output amount of the other asset
// https://github.com/Uniswap/v2-periphery/blob/master/contracts/libraries/UniswapV2Library.sol#L43
func getAmountOut(inputAmount, reserveIn, reserveOut decimal.Decimal) (decimal.Decimal, error) {
	if inputAmount.Cmp(decimal.Zero) <= 0 {
		return decimal.Decimal{}, errInsufficientInputAmount
	}

	if reserveIn.Cmp(decimal.Zero) <= 0 || reserveOut.Cmp(decimal.Zero) <= 0 {
		return decimal.Decimal{}, errInsufficientLiquidity
	}

	inputAmountWithFee := inputAmount.Mul(decimal.New(997, 0))
	numerator := inputAmountWithFee.Mul(reserveOut)
	denominator := reserveIn.Mul(dec1k).Add(inputAmountWithFee)

	return numerator.Div(denominator).RoundDown(0), nil
}
