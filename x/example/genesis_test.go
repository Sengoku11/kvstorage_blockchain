package example_test

import (
	"testing"

	keepertest "example/testutil/keeper"
	"example/testutil/nullify"
	"example/x/example"
	"example/x/example/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		TextStorageList: []types.TextStorage{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ExampleKeeper(t)
	example.InitGenesis(ctx, *k, genesisState)
	got := example.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.TextStorageList, got.TextStorageList)
	// this line is used by starport scaffolding # genesis/test/assert
}
