package keeper

import (
	"context"

	"kvstore/x/kvstore/types"

	"cosmossdk.io/errors"
	"cosmossdk.io/store/prefix"
)

func (k Keeper) Put(ctx context.Context, store prefix.Store, entry types.Entry) error {
	if store.Has([]byte(entry.Key)) {
		return errors.Wrap(types.ErrEntryExist, entry.Key)
	}

	bz, err := k.cdc.Marshal(&types.Entry{
		Value: entry.Value,
	})
	if err != nil {
		return errors.Wrap(err, "Marshal entry failed")
	}

	store.Set([]byte(entry.Key), bz)
	return nil
}

func (k Keeper) Get(ctx context.Context, store prefix.Store, key string) (types.Entry, error) {
	var result types.Entry

	// get can return nil
	if !store.Has([]byte(key)) {
		return result, errors.Wrap(types.ErrEntryExist, key)
	}

	bz := store.Get([]byte(key))
	
	if err := k.cdc.Unmarshal(bz, &result); err != nil {
		return result, err
	}

	return result, nil
}
