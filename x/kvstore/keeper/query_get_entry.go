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

	value := store.Get([]byte(req.Key))
	if value == nil {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetEntryResponse{
		Entry: &types.Entry{
			Key:   req.Key,
			Value: string(value),
		},
	}, nil
}
