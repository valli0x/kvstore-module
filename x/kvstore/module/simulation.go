package kvstore

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"kvstore/testutil/sample"
	kvstoresimulation "kvstore/x/kvstore/simulation"
	"kvstore/x/kvstore/types"
)

// avoid unused import issue
var (
	_ = kvstoresimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgSetEntry = "op_weight_msg_set_entry"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSetEntry int = 100

	opWeightMsgDeleteEntry = "op_weight_msg_delete_entry"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteEntry int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	kvstoreGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&kvstoreGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgSetEntry int
	simState.AppParams.GetOrGenerate(opWeightMsgSetEntry, &weightMsgSetEntry, nil,
		func(_ *rand.Rand) {
			weightMsgSetEntry = defaultWeightMsgSetEntry
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSetEntry,
		kvstoresimulation.SimulateMsgSetEntry(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteEntry int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteEntry, &weightMsgDeleteEntry, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteEntry = defaultWeightMsgDeleteEntry
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteEntry,
		kvstoresimulation.SimulateMsgDeleteEntry(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgSetEntry,
			defaultWeightMsgSetEntry,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				kvstoresimulation.SimulateMsgSetEntry(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteEntry,
			defaultWeightMsgDeleteEntry,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				kvstoresimulation.SimulateMsgDeleteEntry(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
