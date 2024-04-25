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

func (k Keeper) GetEntry(goCtx context.Context, req *types.QueryGetEntryRequest) (*types.QueryGetEntryResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.StoreKey))

	entry, err := k.Get(ctx, store, req.Key)
	if err != nil {
		return &types.QueryGetEntryResponse{}, nil
	}

	return &types.QueryGetEntryResponse{
		Entry: &entry,
	}, nil
}
