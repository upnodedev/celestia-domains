package keeper

import (
	"context"

	"celestia-domain/x/celestiadomain/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k msgServer) SetPrimaryDomain(goCtx context.Context, msg *types.MsgSetPrimaryDomain) (*types.MsgSetPrimaryDomainResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check for ownership and expiration
	parentDomain, err := k.Domain(goCtx, &types.QueryDomainRequest{
		Domain: msg.Domain,
	})
	if err != nil {
		return nil, status.Error(codes.FailedPrecondition, "primary domain doesn't exists or has some error")
	}

	if parentDomain.Owner != msg.Creator {
		return nil, status.Error(codes.PermissionDenied, "primary domain doesn't owned by you")
	}

	if parentDomain.Expiration < uint64(ctx.BlockTime().Unix()) {
		return nil, status.Error(codes.PermissionDenied, "primary domain is expired")
	}

	// Register primary domain
	store := ctx.KVStore(k.storeKey)
	primaryStore := prefix.NewStore(store, []byte("primary-domain"))
	primaryStore.Set([]byte(msg.Creator), []byte(msg.Domain))

	return &types.MsgSetPrimaryDomainResponse{}, nil
}
