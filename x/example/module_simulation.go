package example

import (
	"math/rand"

	"example/testutil/sample"
	examplesimulation "example/x/example/simulation"
	"example/x/example/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = examplesimulation.FindAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
	_ = rand.Rand{}
)

const (
	opWeightMsgCreateTextStorage = "op_weight_msg_text_storage"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateTextStorage int = 100

	opWeightMsgUpdateTextStorage = "op_weight_msg_text_storage"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateTextStorage int = 100

	opWeightMsgDeleteTextStorage = "op_weight_msg_text_storage"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteTextStorage int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	exampleGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		TextStorageList: []types.TextStorage{
			{
				Creator: sample.AccAddress(),
				Index:   "0",
			},
			{
				Creator: sample.AccAddress(),
				Index:   "1",
			},
		},
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&exampleGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// ProposalContents doesn't return any content functions for governance proposals.
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateTextStorage int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateTextStorage, &weightMsgCreateTextStorage, nil,
		func(_ *rand.Rand) {
			weightMsgCreateTextStorage = defaultWeightMsgCreateTextStorage
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateTextStorage,
		examplesimulation.SimulateMsgCreateTextStorage(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateTextStorage int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateTextStorage, &weightMsgUpdateTextStorage, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateTextStorage = defaultWeightMsgUpdateTextStorage
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateTextStorage,
		examplesimulation.SimulateMsgUpdateTextStorage(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteTextStorage int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteTextStorage, &weightMsgDeleteTextStorage, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteTextStorage = defaultWeightMsgDeleteTextStorage
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteTextStorage,
		examplesimulation.SimulateMsgDeleteTextStorage(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateTextStorage,
			defaultWeightMsgCreateTextStorage,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				examplesimulation.SimulateMsgCreateTextStorage(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateTextStorage,
			defaultWeightMsgUpdateTextStorage,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				examplesimulation.SimulateMsgUpdateTextStorage(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteTextStorage,
			defaultWeightMsgDeleteTextStorage,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				examplesimulation.SimulateMsgDeleteTextStorage(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
