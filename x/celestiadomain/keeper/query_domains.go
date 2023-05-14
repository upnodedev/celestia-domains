package keeper

import (
	"context"

	"celestia-domain/x/celestiadomain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Domains(goCtx context.Context, req *types.QueryDomainsRequest) (*types.QueryDomainsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	_ = ctx

	return &types.QueryDomainsResponse{}, nil
}