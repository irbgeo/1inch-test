package pool

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"

	"1inch-test/models"
)

var (
	gasLimit = big.NewInt(10000)
)

type poolProvider struct {
	poolContract      poolContract
	multicallContract multicallContract
}

type poolContract interface {
	Token0CallData() []byte
	Token1CallData() []byte
	GetReservesCallData() []byte
}

type multicallContract interface {
	Multicall(ctx context.Context, call []models.Call) ([]models.CallResult, error)
}

func NewProvider(
	p poolContract,
	m multicallContract,
) *poolProvider {
	return &poolProvider{
		poolContract:      p,
		multicallContract: m,
	}
}

func (s *poolProvider) GetByID(ctx context.Context, poolID common.Address) (models.IPool, error) {
	call := make([]models.Call, 0, 3)

	call = append(call, models.Call{
		Target:   poolID,
		CallData: s.poolContract.Token0CallData(),
		GasLimit: gasLimit,
	})

	call = append(call, models.Call{
		Target:   poolID,
		CallData: s.poolContract.Token1CallData(),
		GasLimit: gasLimit,
	})

	call = append(call, models.Call{
		Target:   poolID,
		CallData: s.poolContract.GetReservesCallData(),
		GasLimit: gasLimit,
	})

	result, err := s.multicallContract.Multicall(ctx, call)
	if err != nil {
		return nil, err
	}

	if len(result) != 3 {
		return nil, errGetPoolInfo
	}

	p := &pool{}

	for i, r := range result {
		if !r.Success {
			return nil, errGetPoolInfo
		}

		switch i {
		case 0:
			p.token0 = common.BytesToAddress(r.ReturnData)
		case 1:
			p.token1 = common.BytesToAddress(r.ReturnData)
		case 2:
			if len(r.ReturnData) < 64 {
				return nil, errGetPoolInfo
			}

			reserve0 := new(big.Int).SetBytes(r.ReturnData[:32])
			reserve1 := new(big.Int).SetBytes(r.ReturnData[32:64])

			p.reserve0 = decimal.NewFromBigInt(reserve0, 0)

			p.reserve1 = decimal.NewFromBigInt(reserve1, 0)

		}
	}

	return p, nil
}
