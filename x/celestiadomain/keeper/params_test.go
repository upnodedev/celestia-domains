package keeper_test

import (
	"testing"

	testkeeper "celestia-domain/testutil/keeper"
	"celestia-domain/x/celestiadomain/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.CelestiadomainKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
