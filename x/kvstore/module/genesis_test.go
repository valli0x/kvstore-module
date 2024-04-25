package kvstore_test

import (
	"testing"

	keepertest "kvstore/testutil/keeper"
	"kvstore/testutil/nullify"
	kvstore "kvstore/x/kvstore/module"
	"kvstore/x/kvstore/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.KvstoreKeeper(t)
	kvstore.InitGenesis(ctx, k, genesisState)
	got := kvstore.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
