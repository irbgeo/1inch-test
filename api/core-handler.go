package api

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"

	"github.com/irbgeo/1inch-test/controller"
)

// GetAmountOut
// @Tags        requests
// @Summary get amount out
// @Description Return outputAmount that corresponding uniswap_v2 pool will return if you try to swap inputAmount of fromToken in poolID
// @Param		fromToken		query	string 		true 	"from token address"	default(0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2)
// @Param		inputAmount		query 	string 		true 	"amount for swapping"	default(1e18)
// @Param		toToken 		query 	string 		true 	"to token address" 		default(0xdac17f958d2ee523a2206206994597c13d831ec7)
// @Param		poolID			query 	string 		true	"pool address"			default(0x0d4a11d5eeaac28ec3f61d100daf4d40471f1852)
// @Success     200	{string}  	string	"amountOut"
// @Failure 	500 {string} 	string	"error description"
// @Failure 	400 {string} 	string	"error description"
// @Router /get-amount-out [get]
func (s *api) getAmountOut(w http.ResponseWriter, r *http.Request) {
	in, err := parseGetAmountOutRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error())) //nolint: errcheck
		return
	}

	out, err := s.controller.GetAmountOut(r.Context(), in)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error())) //nolint: errcheck
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(out.AmountOut.String())) //nolint: errcheck
}

func parseGetAmountOutRequest(r *http.Request) (controller.In, error) {
	query := r.URL.Query()

	poolID := query["poolID"]
	if len(poolID) != 1 {
		return controller.In{}, errPoolIDNumber
	}

	if !common.IsHexAddress(poolID[0]) {
		return controller.In{}, fmt.Errorf("poolID %s is not a hex address", poolID[0])
	}

	fromToken := query["fromToken"]
	if len(fromToken) != 1 {
		return controller.In{}, errFromTokenNumber
	}

	if !common.IsHexAddress(fromToken[0]) {
		return controller.In{}, fmt.Errorf("fromToken %s is not a hex address", fromToken[0])
	}

	inputAmount := query["inputAmount"]
	if len(inputAmount) != 1 {
		return controller.In{}, errInputAmountNumber
	}

	inputAmountValue, err := parseDecimal(inputAmount[0])
	if err != nil {
		return controller.In{}, errParseInputAmount
	}

	toToken := query["toToken"]
	if len(toToken) != 1 {
		return controller.In{}, errToTokenNumber
	}

	if !common.IsHexAddress(toToken[0]) {
		return controller.In{}, fmt.Errorf("toToken %s is not a hex address", toToken[0])
	}

	return controller.In{
		PoolID:      common.HexToAddress(poolID[0]),
		FromToken:   common.HexToAddress(fromToken[0]),
		InputAmount: inputAmountValue,
		ToToken:     common.HexToAddress(toToken[0]),
	}, nil
}

func parseDecimal(src string) (decimal.Decimal, error) {
	src = strings.ToLower(src)
	decimalParts := strings.Split(src, "e")

	if len(decimalParts) == 2 {
		value, err := strconv.ParseInt(decimalParts[0], 10, 64)
		if err != nil {
			return decimal.Decimal{}, err
		}

		exp, err := strconv.ParseInt(decimalParts[1], 10, 32)
		if err != nil {
			return decimal.Decimal{}, err
		}

		return decimal.New(value, int32(exp)), nil
	}

	return decimal.NewFromString(src)
}
