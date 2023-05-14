package keeper

import (
	"context"
	"encoding/json"

	"celestia-domain/x/celestiadomain/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Domain(goCtx context.Context, req *types.QueryDomainRequest) (*types.QueryDomainResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)

	domainStore := prefix.NewStore(store, []byte("domain"))
	if !domainStore.Has([]byte(req.Domain)) {
		return nil, status.Error(codes.NotFound, "domain not found")
	}
	domainJson := domainStore.Get([]byte(req.Domain))

	var domain DomainType
	err := json.Unmarshal([]byte(domainJson), &domain)

	if err != nil {
		return nil, status.Error(codes.Internal, "cannot unmarshal json")
	}

	return &types.QueryDomainResponse{
		Owner:      domain.Owner,
		Domain:     domain.Domain,
		Expiration: domain.Expiration,
	}, nil
}
