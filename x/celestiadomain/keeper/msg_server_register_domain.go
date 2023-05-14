package keeper

import (
	"context"

	"celestia-domain/x/celestiadomain/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const ONE_YEAR = 31536000

func (k msgServer) RegisterDomain(goCtx context.Context, msg *types.MsgRegisterDomain) (*types.MsgRegisterDomainResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgRegisterDomainResponse{}, nil
}
