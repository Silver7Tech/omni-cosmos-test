package keeper

import (
	"context"

	"test/x/test/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func BoolToString(b bool) string {
	if b {
		return "true"
	}
	return "false"
}

func (k Keeper) ShowStorage(goCtx context.Context, req *types.QueryShowStorageRequest) (*types.QueryShowStorageResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// omni, found := k.GetOmni(ctx, req.Id)
	// ctx.Logger().Info(BoolToString(found))
	// if !found {
	// 	return nil, sdkerrors.ErrKeyNotFound
	// }
	// _ = omni
	return &types.QueryShowStorageResponse{Omni: PrefixStoreToString(ctx, ctx.KVStore(k.storeKey), types.KeyPrefix(types.OmniKey))}, nil
}
