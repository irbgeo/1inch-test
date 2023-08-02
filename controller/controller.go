package controller

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
)

type controller struct {
	poolProvider poolProvider
}

type poolProvider interface {
	GetByID(ctx context.Context, poolID common.Address) (IPool, error)
}

type IPool interface {
	GetToken0() common.Address
	GetToken1() common.Address
	GetAmountOut(fromToken common.Address, inputAmount decimal.Decimal) (decimal.Decimal, error)
}

func New(poolProvider poolProvider) *controller {
	return &controller{poolProvider: poolProvider}
}

func (s *controller) GetAmountOut(ctx context.Context, in In) (*Out, error) {
	pool, err := s.poolProvider.GetByID(ctx, in.PoolID)
	if err != nil {
		return nil, fmt.Errorf("failed get pool %s: %w", in.PoolID, err)
	}

	if (pool.GetToken0() != in.FromToken && pool.GetToken1() != in.ToToken) &&
		(pool.GetToken0() != in.ToToken && pool.GetToken1() != in.FromToken) {
		return nil, fmt.Errorf("pool %s does not correspond to the exchanged tokens want %s/%s have %s/%s",
			in.PoolID, in.FromToken, in.ToToken, pool.GetToken0(), pool.GetToken1())
	}

	out := new(Out)

	out.AmountOut, err = pool.GetAmountOut(in.FromToken, in.InputAmount)
	if err != nil {
		return nil, fmt.Errorf("get amount out: %w", err)
	}

	return out, nil
}
