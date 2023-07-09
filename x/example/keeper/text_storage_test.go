package keeper_test

import (
	"strconv"
	"testing"

	keepertest "example/testutil/keeper"
	"example/testutil/nullify"
	"example/x/example/keeper"
	"example/x/example/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNTextStorage(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.TextStorage {
	items := make([]types.TextStorage, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetTextStorage(ctx, items[i])
	}
	return items
}

func TestTextStorageGet(t *testing.T) {
	keeper, ctx := keepertest.ExampleKeeper(t)
	items := createNTextStorage(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetTextStorage(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestTextStorageRemove(t *testing.T) {
	keeper, ctx := keepertest.ExampleKeeper(t)
	items := createNTextStorage(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveTextStorage(ctx,
			item.Index,
		)
		_, found := keeper.GetTextStorage(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestTextStorageGetAll(t *testing.T) {
	keeper, ctx := keepertest.ExampleKeeper(t)
	items := createNTextStorage(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllTextStorage(ctx)),
	)
}
