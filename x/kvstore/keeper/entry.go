package keeper

import (
	"context"

	"kvstore/x/kvstore/types"

	"cosmossdk.io/errors"
	"cosmossdk.io/store/prefix"
)

func (k Keeper) Put(ctx context.Context, store prefix.Store, entry types.Entry) error {
	keyb := types.KeyPrefix(entry.Key)

	if store.Has(keyb) {
		return errors.Wrap(types.ErrEntryExist, entry.Key)
	}

	bz, err := k.cdc.Marshal(&types.Entry{
		Value: entry.Value,
	})
	if err != nil {
		return errors.Wrap(types.ErrEntryNotExist, entry.Key)
	}

	store.Set(keyb, bz)
	return nil
}

func (k Keeper) Get(ctx context.Context, store prefix.Store, key string) (types.Entry, error) {
	var result types.Entry
	keyb := types.KeyPrefix(key)

	// get can return nil
	if !store.Has(keyb) {
		return result, errors.Wrap(types.ErrEntryExist, key)
	}

	bz := store.Get(keyb)

	if err := k.cdc.Unmarshal(bz, &result); err != nil {
		return result, err
	}

	return result, nil
}
