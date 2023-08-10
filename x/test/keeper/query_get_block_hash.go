package keeper

import (
	"context"

	"test/x/ethservice"
	"test/x/test/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetBlockHash(goCtx context.Context, req *types.QueryGetBlockHashRequest) (*types.QueryGetBlockHashResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	ethService, err := ethservice.NewEthService("https://mainnet.infura.io/v3/46d7e5b569734fad998bf16a3c13b3e4")
	if err != nil {
		return nil, err
	}
	blockHash, err := ethService.GetBlockHashByNumber(req.BlockNumber)
	if err != nil {
		return nil, err
	}
	// TODO: Process the query
	_ = ctx

	return &types.QueryGetBlockHashResponse{BlockHash: blockHash}, nil
}
