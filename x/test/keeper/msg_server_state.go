package keeper

import (
	"context"

	"test/x/test/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) State(goCtx context.Context, msg *types.MsgState) (*types.MsgStateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgStateResponse{}, nil
}
