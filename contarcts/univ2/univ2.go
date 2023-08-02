package univ2

import (
	_ "embed"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

var (
	//go:embed abi/univ2.json
	univ2ABI string
)

type univ2 struct {
	abi abi.ABI

	token0Calldata      []byte
	token1Calldata      []byte
	getReservesCallData []byte
}

func NewContract() (*univ2, error) {
	var (
		u   = new(univ2)
		err error
	)

	u.abi, err = abi.JSON(strings.NewReader(univ2ABI))
	if err != nil {
		return nil, err
	}

	u.token0Calldata, err = u.abi.Pack("token0")
	if err != nil {
		return nil, err
	}

	u.token1Calldata, err = u.abi.Pack("token1")
	if err != nil {
		return nil, err
	}

	u.getReservesCallData, err = u.abi.Pack("getReserves")
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (s *univ2) Token0CallData() []byte {
	return s.token0Calldata
}

func (s *univ2) Token1CallData() []byte {
	return s.token1Calldata
}

func (s *univ2) GetReservesCallData() []byte {
	return s.getReservesCallData
}
