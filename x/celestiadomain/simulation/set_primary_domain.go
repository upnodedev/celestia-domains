package simulation

import (
	"math/rand"

	"celestia-domain/x/celestiadomain/keeper"
	"celestia-domain/x/celestiadomain/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgSetPrimaryDomain(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgSetPrimaryDomain{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the SetPrimaryDomain simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "SetPrimaryDomain simulation not implemented"), nil, nil
	}
}
