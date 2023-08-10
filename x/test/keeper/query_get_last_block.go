package keeper

import (
	"context"
	"fmt"

	"test/x/ethservice"
	"test/x/test/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetLastBlock(goCtx context.Context, req *types.QueryGetLastBlockRequest) (*types.QueryGetLastBlockResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	ethService, err := ethservice.NewEthService("https://mainnet.infura.io/v3/46d7e5b569734fad998bf16a3c13b3e4")
	if err != nil {
		return nil, err
	}
	lastblock, error := ethService.GetLatestBlockNumber()
	if error != nil {
		return nil, error
	}
	// TODO: Process the query
	_ = ctx

	return &types.QueryGetLastBlockResponse{BlockNumber: fmt.Sprintf("%d", lastblock)}, nil
}
