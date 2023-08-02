package api

import (
	"fmt"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"

	"1inch-test/models"
)

// GetAmountOut godoc
// @Summary get amount out 
// @Description Return outputAmount that corresponding uniswap_v2 pool  will return if you try to swap inputAmount of  fromToken
// @Param	0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2	query	string 		true 	"fromToken"
// @Param	1000000000000000000							query 	string 		true 	"amount for swapping"
// @Param	0xdac17f958d2ee523a2206206994597c13d831ec7 	query 	string 		true 	"toToken"
// @Param	0x0d4a11d5eeaac28ec3f61d100daf4d40471f1852  query 	string 		true	"poolID"
// @Success     200	{string}  	string	"amountOut"
// @Failure 	500 {string} 	string	"error description"
// @Failure 	404 {string} 	string	"error description"
// @Router /get-amount-out [get]
func (s *api) getAmountOut(w http.ResponseWriter, r *http.Request) {
	in, err := parseGetAmountOutRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error())) //nolint: errcheck
		return
	}

	out, err := s.core.GetAmountOut(r.Context(), in)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error())) //nolint: errcheck
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(out.AmountOut.String())) //nolint: errcheck
}

func parseGetAmountOutRequest(r *http.Request) (models.In, error) {
	query := r.URL.Query()

	poolID := query["poolID"]
	if len(poolID) != 1 {
		return models.In{}, errPoolIDNumber
	}

	if !common.IsHexAddress(poolID[0]) {
		return models.In{}, fmt.Errorf("poolID %s is not a hex address", poolID[0])
	}

	fromToken := query["fromToken"]
	if len(fromToken) != 1 {
		return models.In{}, errFromTokenNumber
	}

	if !common.IsHexAddress(fromToken[0]) {
		return models.In{}, fmt.Errorf("fromToken %s is not a hex address", fromToken[0])
	}

	inputAmount := query["inputAmount"]
	if len(inputAmount) != 1 {
		return models.In{}, errInputAmountNumber
	}

	inputAmountValue, err := decimal.NewFromString(inputAmount[0])
	if err != nil {
		return models.In{}, errParseInputAmount
	}

	toToken := query["toToken"]
	if len(toToken) != 1 {
		return models.In{}, errToTokenNumber
	}

	if !common.IsHexAddress(toToken[0]) {
		return models.In{}, fmt.Errorf("toToken %s is not a hex address", toToken[0])
	}

	return models.In{
		PoolID:      common.HexToAddress(poolID[0]),
		FromToken:   common.HexToAddress(fromToken[0]),
		InputAmount: inputAmountValue,
		ToToken:     common.HexToAddress(toToken[0]),
	}, nil
}
