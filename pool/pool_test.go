package pool

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

type testCase struct {
	amountIn          decimal.Decimal
	reserveIn         decimal.Decimal
	reserveOut        decimal.Decimal
	expectedAmountOut decimal.Decimal
	expectedError     error
}

var (
	reserveInTest, _ = decimal.NewFromString("17278533068023270614806")
	reserveOut, _    = decimal.NewFromString("32013064463681")

	testCases = []testCase{
		// https://github.com/Uniswap/v2-periphery/blob/master/test/UniswapV2Router02.spec.ts#L54
		{
			amountIn:          decimal.New(2, 0),
			reserveIn:         decimal.New(1, 2),
			reserveOut:        decimal.New(1, 2),
			expectedAmountOut: decimal.New(1, 0),
			expectedError:     nil,
		},
		{
			reserveIn:     decimal.New(1, 2),
			reserveOut:    decimal.New(1, 2),
			expectedError: errInsufficientInputAmount,
		},
		{
			amountIn:      decimal.New(2, 0),
			reserveIn:     decimal.New(1, 2),
			expectedError: errInsufficientLiquidity,
		},
		{
			amountIn:      decimal.New(2, 0),
			reserveOut:    decimal.New(1, 2),
			expectedError: errInsufficientLiquidity,
		},
		{
			amountIn:          decimal.New(1, 18),
			reserveIn:         reserveInTest,
			reserveOut:        reserveOut,
			expectedAmountOut: decimal.New(1847100305, 0),
			expectedError:     nil,
		},
	}
)

func TestGetAmountOut(t *testing.T) {
	for _, testCase := range testCases {
		actualAmountOut, actualError := getAmountOut(testCase.amountIn, testCase.reserveIn, testCase.reserveOut)
		require.Equal(t, testCase.expectedError, actualError)
		require.Equal(t, testCase.expectedAmountOut, actualAmountOut)
	}
}
