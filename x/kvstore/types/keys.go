package types

const (
	// ModuleName defines the module name
	ModuleName = "kvstore"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_kvstore"
)

var (
	ParamsKey = []byte("p_kvstore")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
