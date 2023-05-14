package keeper

import (
	"context"
	"strings"

	"celestia-domain/x/celestiadomain/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) PrimaryDomain(goCtx context.Context, req *types.QueryPrimaryDomainRequest) (*types.QueryPrimaryDomainResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)

	primaryStore := prefix.NewStore(store, []byte("primary-domain"))
	if !primaryStore.Has([]byte(req.Owner)) {
		return nil, status.Error(codes.NotFound, "primary domain not set")
	}

	// Get domain object
	domainName := string(primaryStore.Get([]byte(req.Owner)))
	domain, err := k.Domain(goCtx, &types.QueryDomainRequest{
		Domain: domainName,
	})
	if err != nil {
		return nil, status.Error(codes.NotFound, "domain name not found or error")
	}

	// Build parent
	parts := strings.Split(domainName, ".")
	parent := strings.Join(parts[1:], ".")

	return &types.QueryPrimaryDomainResponse{
		Domain: &types.Domain{
			Owner:      domain.Owner,
			Domain:     domain.Domain,
			Parent:     parent,
			Expiration: domain.Expiration,
		},
	}, nil
}
