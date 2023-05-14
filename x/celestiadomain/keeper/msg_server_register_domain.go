package keeper

import (
	"context"
	"encoding/binary"
	"encoding/json"
	"strings"

	"celestia-domain/x/celestiadomain/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetOwnedDomainsCount(ctx sdk.Context, owner string) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte("owned-domains-count"))

	// Use owner as count key
	byteKey := []byte(owner)

	// Get the value of the count
	bz := store.Get(byteKey)

	// Return zero if the count value is not found (for example, it's the first domain)
	if bz == nil {
		return 0
	}

	// Convert the count into a uint64
	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) SetOwnedDomainsCount(ctx sdk.Context, owner string, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte("owned-domains-count"))

	// Use owner as count key
	byteKey := []byte(owner)

	// Convert count from uint64 to string and get bytes
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)

	// Set the value of domain-count- to count
	store.Set(byteKey, bz)
}

func (k Keeper) AppendOwnedDomains(ctx sdk.Context, owner string, domain string) uint64 {
	// Get the current number of domains in the store
	count := k.GetOwnedDomainsCount(ctx, owner)

	// Get the store
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte("owned-domains-"+owner))

	// Convert the domain ID into bytes
	byteKey := make([]byte, 8)
	binary.BigEndian.PutUint64(byteKey, count)

	// Map count -> domain (Append)
	store.Set(byteKey, []byte(domain))

	// Update the domain count
	k.SetOwnedDomainsCount(ctx, owner, count+1)
	return count
}

func (k msgServer) RegisterDomain(goCtx context.Context, msg *types.MsgRegisterDomain) (*types.MsgRegisterDomainResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate domain validity
	_, err := k.Domain(goCtx, &types.QueryDomainRequest{
		Domain: msg.Domain,
	})
	if err == nil {
		return nil, status.Error(codes.FailedPrecondition, "domain has already taken")
	}

	parts := strings.Split(msg.Domain, ".")
	if len(parts) < 2 {
		return nil, status.Error(codes.FailedPrecondition, "domain must have at least two parts")
	}
	if parts[len(parts)-1] != "tia" {
		return nil, status.Error(codes.FailedPrecondition, "only .tia tld supported")
	}
	parent := strings.Join(parts[1:], ".")

	// Check parent ownership if registering a new subdomain
	if len(parts) > 2 {
		parentDomain, err := k.Domain(goCtx, &types.QueryDomainRequest{
			Domain: parent,
		})
		if err != nil {
			return nil, status.Error(codes.FailedPrecondition, "parent domain doesn't exists or has some error")
		}

		if parentDomain.Owner != msg.Creator {
			return nil, status.Error(codes.PermissionDenied, "parent domain doesn't owned by you")
		}

		if parentDomain.Expiration < uint64(ctx.BlockTime().Unix()) {
			return nil, status.Error(codes.PermissionDenied, "parent domain is expired")
		}
	}

	// Define domain object
	expiration := ctx.BlockTime().AddDate(1, 0, 0)
	domain := DomainType{
		Owner:      msg.Creator,
		Domain:     msg.Domain,
		Parent:     parent,
		Expiration: uint64(expiration.Unix()),
	}

	// Register domain to domain store
	store := ctx.KVStore(k.storeKey)
	domainStore := prefix.NewStore(store, []byte("domain"))
	domainJson, err := json.Marshal(domain)
	if err != nil {
		return nil, status.Error(codes.Internal, "cannot unmarshal json")
	}
	domainStore.Set([]byte(msg.Domain), domainJson)

	// Register domain to owner domain list store
	k.AppendOwnedDomains(ctx, msg.Creator, msg.Domain)

	return &types.MsgRegisterDomainResponse{
		Owner:      msg.Creator,
		Domain:     msg.Domain,
		Expiration: domain.Expiration,
	}, nil
}
