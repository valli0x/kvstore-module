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

	store.Set([]byte(msg.Key), []byte(msg.Value))

	return &types.MsgSetEntryResponse{}, nil
}
