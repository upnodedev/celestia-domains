package celestiadomain

import (
	"math/rand"

	"celestia-domain/testutil/sample"
	celestiadomainsimulation "celestia-domain/x/celestiadomain/simulation"
	"celestia-domain/x/celestiadomain/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = celestiadomainsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgRegisterDomain = "op_weight_msg_register_domain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRegisterDomain int = 100

	opWeightMsgSetPrimaryDomain = "op_weight_msg_set_primary_domain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSetPrimaryDomain int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	celestiadomainGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&celestiadomainGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgRegisterDomain int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRegisterDomain, &weightMsgRegisterDomain, nil,
		func(_ *rand.Rand) {
			weightMsgRegisterDomain = defaultWeightMsgRegisterDomain
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRegisterDomain,
		celestiadomainsimulation.SimulateMsgRegisterDomain(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSetPrimaryDomain int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSetPrimaryDomain, &weightMsgSetPrimaryDomain, nil,
		func(_ *rand.Rand) {
			weightMsgSetPrimaryDomain = defaultWeightMsgSetPrimaryDomain
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSetPrimaryDomain,
		celestiadomainsimulation.SimulateMsgSetPrimaryDomain(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
