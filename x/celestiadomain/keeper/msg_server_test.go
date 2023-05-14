package keeper_test

import (
	"context"
	"testing"

	keepertest "celestia-domain/testutil/keeper"
	"celestia-domain/x/celestiadomain/keeper"
	"celestia-domain/x/celestiadomain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.CelestiadomainKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
