package multicall

import (
	"context"
	_ "embed"
	"log"
	"strings"

	"github.com/irbgeo/1inch-test/pool"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	//go:embed abi/multicall.json
	multicallABI string

	multicallAddressContract = common.HexToAddress("0x1F98415757620B543A52E61c46B32eB19261F984")
)

type multicall struct {
	abi abi.ABI

	client *ethclient.Client
}

func NewContract(providerURL string) (*multicall, error) {
	var (
		m   = new(multicall)
		err error
	)

	m.abi, err = abi.JSON(strings.NewReader(multicallABI))
	if err != nil {
		return nil, err
	}

	m.client, err = ethclient.Dial(providerURL)
	if err != nil {
		log.Fatal(err)
	}

	return m, nil
}

func (s *multicall) Multicall(ctx context.Context, call []pool.Call) ([]pool.CallResult, error) {
	calldata, err := s.abi.Pack("multicall", call)
	if err != nil {
		return nil, err
	}

	msg := ethereum.CallMsg{
		To:   &multicallAddressContract,
		Data: calldata,
	}

	data, err := s.client.CallContract(ctx, msg, nil)
	if err != nil {
		return nil, err
	}

	var returnData multicallResponse

	err = s.abi.UnpackIntoInterface(&returnData, "multicall", data)

	out := make([]pool.CallResult, 0, len(returnData.ReturnData))

	for _, r := range returnData.ReturnData {
		out = append(out, pool.CallResult{
			Success:    r.Success,
			ReturnData: r.ReturnData,
		})
	}

	return out, err
}
