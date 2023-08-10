package ethservice

import (
	"strconv"

	"github.com/ethereum/go-ethereum/rpc"
)

type EthService struct {
	client *rpc.Client
}

func NewEthService(rpcURL string) (*EthService, error) {
	client, err := rpc.Dial(rpcURL)
	if err != nil {
		return nil, err
	}
	return &EthService{client: client}, nil
}

// Now you can add methods for various Ethereum RPC interactions you need

func (es *EthService) GetLatestBlockNumber() (uint64, error) {
	var result string
	err := es.client.Call(&result, "eth_blockNumber")
	if err != nil {
		return 0, err
	}
	// Convert result (in hex) to uint64
	blockNumber, err := strconv.ParseUint(result[2:], 16, 64)
	return blockNumber, err
}
