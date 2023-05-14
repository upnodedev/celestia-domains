package keeper

import (
	"context"

	"celestia-domain/x/celestiadomain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) PrimaryDomain(goCtx context.Context, req *types.QueryPrimaryDomainRequest) (*types.QueryPrimaryDomainResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	_ = ctx

	return &types.QueryPrimaryDomainResponse{}, nil
}
