package keeper

import (
	"context"

	"test/x/ethservice"
	"test/x/test/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetStorage(goCtx context.Context, req *types.QueryGetStorageRequest) (*types.QueryGetStorageResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ethService, err := ethservice.NewEthService("https://mainnet.infura.io/v3/46d7e5b569734fad998bf16a3c13b3e4")
	if err != nil {
		return nil, err
	}
	storage, err := ethService.GetStorageAt(req.Address, req.Position, req.BlockTag)
	if err != nil {
		return nil, err
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	_ = ctx

	return &types.QueryGetStorageResponse{Storage: storage}, nil
}
