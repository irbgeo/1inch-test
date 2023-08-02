package api

import (
	"fmt"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"

	"1inch-test/models"
)

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

	tokenIn := query["tokenIn"]
	if len(tokenIn) != 1 {
		return models.In{}, errTokenInNumber
	}

	if !common.IsHexAddress(tokenIn[0]) {
		return models.In{}, fmt.Errorf("tokenIn %s is not a hex address", tokenIn[0])
	}

	amountIn := query["amountIn"]
	if len(amountIn) != 1 {
		return models.In{}, errAmountInNumber
	}

	amountInValue, err := decimal.NewFromString(amountIn[0])
	if err != nil {
		return models.In{}, errParseAmountIn
	}

	tokenOut := query["tokenOut"]
	if len(tokenOut) != 1 {
		return models.In{}, errTokenOutNumber
	}

	if !common.IsHexAddress(tokenOut[0]) {
		return models.In{}, fmt.Errorf("tokenOut %s is not a hex address", tokenOut[0])
	}

	return models.In{
		PoolID:   common.HexToAddress(poolID[0]),
		TokenIn:  common.HexToAddress(tokenIn[0]),
		AmountIn: amountInValue,
		TokenOut: common.HexToAddress(tokenOut[0]),
	}, nil
}
