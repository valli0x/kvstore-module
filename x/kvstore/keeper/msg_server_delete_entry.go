package keeper

import (
	"context"

	"kvstore/x/kvstore/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) DeleteEntry(goCtx context.Context, msg *types.MsgDeleteEntry) (*types.MsgDeleteEntryResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.StoreKey))

	if err := k.Delete(ctx, store, msg.Key); err != nil {
		return &types.MsgDeleteEntryResponse{}, err
	}

	return &types.MsgDeleteEntryResponse{}, nil
}
