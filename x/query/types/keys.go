package types

const (
	// ModuleName defines the module name
	ModuleName = "query"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_query"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
