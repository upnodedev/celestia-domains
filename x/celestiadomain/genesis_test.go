package celestiadomain_test

import (
	"testing"

	keepertest "celestia-domain/testutil/keeper"
	"celestia-domain/testutil/nullify"
	"celestia-domain/x/celestiadomain"
	"celestia-domain/x/celestiadomain/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CelestiadomainKeeper(t)
	celestiadomain.InitGenesis(ctx, *k, genesisState)
	got := celestiadomain.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
