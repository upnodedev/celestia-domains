package simulation

import (
	"math/rand"

	"celestia-domain/x/celestiadomain/keeper"
	"celestia-domain/x/celestiadomain/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgRegisterDomain(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgRegisterDomain{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the RegisterDomain simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "RegisterDomain simulation not implemented"), nil, nil
	}
}
