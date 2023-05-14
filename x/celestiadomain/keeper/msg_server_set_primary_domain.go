package keeper

import (
	"context"

	"celestia-domain/x/celestiadomain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SetPrimaryDomain(goCtx context.Context, msg *types.MsgSetPrimaryDomain) (*types.MsgSetPrimaryDomainResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSetPrimaryDomainResponse{}, nil
}
