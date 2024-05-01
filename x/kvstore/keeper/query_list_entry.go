package keeper

import (
	"context"

	"kvstore/x/kvstore/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ListEntry(goCtx context.Context, req *types.QueryListEntryRequest) (*types.QueryListEntryResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.StoreKey))

	entries, err := k.List(ctx, store, req.Key)
	if err != nil {
		return &types.QueryListEntryResponse{}, err
	}

	return &types.QueryListEntryResponse{
		Entry: entries,
	}, nil
}
