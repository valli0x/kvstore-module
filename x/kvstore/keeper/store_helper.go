package keeper

import (
	"cosmossdk.io/store"
	"cosmossdk.io/store/prefix"
	stortypes "cosmossdk.io/store/types"
	db "github.com/cosmos/cosmos-db"
)

/*
	Iterator
*/

func GatherValuesFromStorePrefix[T any](storeObj store.KVStore, prefix []byte, parseValue func([]byte) (T, error)) ([]T, error) {
	// Replace a callback with the one that takes both key and value
	// but ignores the key.
	parseOnlyValue := func(_ []byte, value []byte) (T, error) {
		return parseValue(value)
	}
	return GatherValuesFromStorePrefixWithKeyParser(storeObj, prefix, parseOnlyValue)
}

func GatherValuesFromStorePrefixWithKeyParser[T any](storeObj store.KVStore, prefix []byte, parse func(key []byte, value []byte) (T, error)) ([]T, error) {
	iterator := stortypes.KVStorePrefixIterator(storeObj, prefix)
	defer iterator.Close()
	return gatherValuesFromIteratorWithKeyParser(iterator, parse, noStopFn)
}

func gatherValuesFromIteratorWithKeyParser[T any](iterator db.Iterator, parse func(key []byte, value []byte) (T, error), stopFn func([]byte) bool) ([]T, error) {
	values := []T{}
	for ; iterator.Valid(); iterator.Next() {
		if stopFn(iterator.Key()) {
			break
		}
		val, err := parse(iterator.Key(), iterator.Value())
		if err != nil {
			return nil, err
		}
		values = append(values, val)
	}
	return values, nil
}

func noStopFn([]byte) bool {
	return false
}

/*
	Delete entries with prefix
*/

// DeleteAllKeysFromPrefix deletes all store records that contains the given prefixKey.
func DeleteAllKeysFromPrefix(store store.KVStore, prefixKey []byte) error {
	prefixStore := prefix.NewStore(store, prefixKey)
	iter := prefixStore.Iterator(nil, nil)
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		prefixStore.Delete(iter.Key())
	}
	return nil
}
