package keeper

import (
	"context"

	"kvstore/x/kvstore/types"

	"cosmossdk.io/errors"
	"cosmossdk.io/store/prefix"
)

// osmosis example https://github.com/osmosis-labs/osmosis/blob/main/osmoutils/store_helper.go#L154

func (k Keeper) Put(ctx context.Context, store prefix.Store, entry types.Entry) error {
	keyb := types.KeyPrefix(entry.Key)

	if store.Has(keyb) {
		return errors.Wrap(types.ErrEntryExist, entry.Key)
	}

	bz, err := k.cdc.Marshal(&types.Entry{
		Value: entry.Value,
	})
	if err != nil {
		return errors.Wrap(err, entry.Key)
	}

	store.Set(keyb, bz)
	return nil
}

func (k Keeper) Get(ctx context.Context, store prefix.Store, key string) (types.Entry, error) {
	var result types.Entry
	keyb := types.KeyPrefix(key)

	// get can return nil
	if !store.Has(keyb) {
		return result, errors.Wrap(types.ErrEntryNotExist, key)
	}

	bz := store.Get(keyb)

	if err := k.cdc.Unmarshal(bz, &result); err != nil {
		return result, err
	}

	return result, nil
}

func (k Keeper) Delete(ctx context.Context, store prefix.Store, key string) error {
	return DeleteAllKeysFromPrefix(store, types.KeyPrefix(key))
}

func (k Keeper) List(ctx context.Context, store prefix.Store, prefixKey string) ([]*types.Entry, error) {
	// parsing entry func
	parseValue := func(value []byte) (*types.Entry, error) {
		entry := &types.Entry{}
		if err := k.cdc.Unmarshal(value, entry); err != nil {
			return &types.Entry{}, err
		}
		return entry, nil
	}

	return GatherValuesFromStorePrefix(store, types.KeyPrefix(prefixKey), parseValue)
}
