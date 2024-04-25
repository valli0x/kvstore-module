package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/kvstore module sentinel errors
var (
	ErrEntryExist    = sdkerrors.Register(ModuleName, 3, "entry already exist")
	ErrInvalidSigner = sdkerrors.Register(ModuleName, 1100, "expected gov account as only signer for proposal message")
	ErrSample        = sdkerrors.Register(ModuleName, 1101, "sample error")
)
