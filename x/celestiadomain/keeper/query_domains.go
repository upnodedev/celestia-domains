package keeper

import (
	"context"
	"encoding/json"

	"celestia-domain/x/celestiadomain/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Domains(goCtx context.Context, req *types.QueryDomainsRequest) (*types.QueryDomainsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get store
	store := ctx.KVStore(k.storeKey)
	domainStore := prefix.NewStore(store, []byte("owned-domains-"+req.Owner))

	// Paginate the recipes store based on PageRequest
	var domains []*types.Domain

	pageRes, err := query.Paginate(domainStore, req.Pagination, func(key []byte, value []byte) error {
		var domain DomainType
		if err := json.Unmarshal(value, &domain); err != nil {
			return err
		}

		domainConverted := types.Domain{
			Owner:      domain.Owner,
			Domain:     domain.Domain,
			Parent:     domain.Parent,
			Expiration: domain.Expiration,
		}

		domains = append(domains, &domainConverted)

		return nil
	})

	// Throw an error if pagination failed
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryDomainsResponse{
		Domains:    domains,
		Pagination: pageRes,
	}, nil
}
