package keeper

import (
	"celestia-domain/x/celestiadomain/types"
)

var _ types.QueryServer = Keeper{}
