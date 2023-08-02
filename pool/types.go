package pool

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type Call struct {
	Target   common.Address
	CallData []byte
	GasLimit *big.Int
}

type CallResult struct {
	Success    bool
	ReturnData []byte
}
