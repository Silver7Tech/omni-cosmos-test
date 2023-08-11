package keeper

import (
	"context"

	"test/x/test/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ListStorage(goCtx context.Context, req *types.QueryListStorageRequest) (*types.QueryListStorageResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var omnis []types.Omni
	ctx := sdk.UnwrapSDKContext(goCtx)
	store := ctx.KVStore(k.storeKey)
	omniStore := prefix.NewStore(store, types.KeyPrefix(types.OmniKey))

	pageRes, err := query.Paginate(omniStore, req.Pagination, func(key []byte, value []byte) error {
		var omni types.Omni
		if err := k.cdc.Unmarshal(value, &omni); err != nil {
			return err
		}
		omnis = append(omnis, omni)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryListStorageResponse{Omni: omnis, Pagination: pageRes}, nil
}
