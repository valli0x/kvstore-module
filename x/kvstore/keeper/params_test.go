package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "kvstore/testutil/keeper"
	"kvstore/x/kvstore/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.KvstoreKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
