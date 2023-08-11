package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgState = "state"

var _ sdk.Msg = &MsgState{}

func NewMsgState(creator string, address string, position string, blocktag string, storage string) *MsgState {
	return &MsgState{
		Creator:  creator,
		Address:  address,
		Position: position,
		Blocktag: blocktag,
		Storage:  storage,
	}
}

func (msg *MsgState) Route() string {
	return RouterKey
}

func (msg *MsgState) Type() string {
	return TypeMsgState
}

func (msg *MsgState) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgState) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgState) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
