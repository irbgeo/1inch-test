package multicall

import "math/big"

type multicallResponse struct {
	BlockNumber *big.Int
	ReturnData  []returnData
}

type returnData struct {
	Success    bool
	GasUsed    *big.Int
	ReturnData []byte
}
