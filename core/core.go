package core

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"

	"1inch-test/models"
)

type core struct {
	poolProvider poolProvider
}

type poolProvider interface {
	GetByID(ctx context.Context, poolID common.Address) (models.IPool, error)
}

func New(poolProvider poolProvider) *core {
	return &core{poolProvider: poolProvider}
}

func (s *core) GetAmountOut(ctx context.Context, in models.In) (*models.Out, error) {
	pool, err := s.poolProvider.GetByID(ctx, in.PoolID)
	if err != nil {
		return nil, fmt.Errorf("failed get pool %s: %w", in.PoolID, err)
	}

	if (pool.GetToken0() != in.TokenIn && pool.GetToken1() != in.TokenOut) &&
		(pool.GetToken0() != in.TokenOut && pool.GetToken1() != in.TokenIn) {
		return nil, fmt.Errorf("pool %s does not correspond to the exchanged tokens want %s/%s have %s/%s",
			in.PoolID, in.TokenIn, in.TokenOut, pool.GetToken0(), pool.GetToken1())
	}

	out := new(models.Out)

	out.AmountOut, err = pool.GetAmountOut(in.TokenIn, in.AmountIn)
	if err != nil {
		return nil, fmt.Errorf("get amount out: %w", err)
	}

	return out, nil
}
