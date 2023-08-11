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

func PrefixStoreToString(ctx sdk.Context, store sdk.KVStore, prefix []byte) string {
	iterator := sdk.KVStorePrefixIterator(store, prefix)
	defer iterator.Close()

	var output string
	for ; iterator.Valid(); iterator.Next() {
		key := string(iterator.Key()[len(prefix):]) // Remove the prefix from the key
		value := string(iterator.Value())
		output += fmt.Sprintf("Key: %s, Value: %s\n", key, value)
	}

	return output
}

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
	var result = types.Omni{
		Address:  req.Address,
		Position: req.Position,
		Blocktag: req.BlockTag,
		Storage:  storage,
	}
	count := k.AppendData(ctx, result)
	_ = count
	return &types.QueryGetStorageResponse{Count: PrefixStoreToString(ctx, ctx.KVStore(k.storeKey), types.KeyPrefix(types.OmniKey))}, nil
}
