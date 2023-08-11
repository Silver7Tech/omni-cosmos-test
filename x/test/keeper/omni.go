package keeper

import (
	"encoding/binary"
	"fmt"
	"test/x/test/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) AppendData(ctx sdk.Context, omni types.Omni) prefix.Store {
	count := k.GetDataCount(ctx)
	omni.Id = count
	fmt.Println("Omni Storage Count", count)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OmniKey))
	appendedValue := k.cdc.MustMarshal(&omni)
	store.Set(GetOmniIDBytes(omni.Id), appendedValue)
	k.SetDataCount(ctx, count+1)
	return store
}

func (k Keeper) GetDataCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.OmniCountKey)
	bz := store.Get(byteKey)
	if bz == nil {
		return 0
	}
	return binary.BigEndian.Uint64(bz)
}

func GetOmniIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

func (k Keeper) SetDataCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.OmniCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

func (k Keeper) GetOmni(ctx sdk.Context, id uint64) (val types.Omni, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OmniKey))
	b := store.Get(GetOmniIDBytes(id))
	if b == nil {
		return val, false
	}

	ctx.Logger().Info(string(b))
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}
