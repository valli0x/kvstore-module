package keeper

import (
	"context"

	"kvstore/x/kvstore/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SetEntry(goCtx context.Context, msg *types.MsgSetEntry) (*types.MsgSetEntryResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.StoreKey))

	if err := k.Put(ctx, store, types.Entry{
		Key: msg.Key,
		Value: msg.Value,
	}); err != nil {
		return &types.MsgSetEntryResponse{}, err
	}

	return &types.MsgSetEntryResponse{}, nil
}
