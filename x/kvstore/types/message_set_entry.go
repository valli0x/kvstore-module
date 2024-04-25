package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetEntry{}

func NewMsgSetEntry(creator string, key string, value string) *MsgSetEntry {
	return &MsgSetEntry{
		Creator: creator,
		Key:     key,
		Value:   value,
	}
}

func (msg *MsgSetEntry) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
