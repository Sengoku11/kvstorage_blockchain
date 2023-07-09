package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "example/testutil/keeper"
	"example/testutil/nullify"
	"example/x/example/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestTextStorageQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.ExampleKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNTextStorage(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetTextStorageRequest
		response *types.QueryGetTextStorageResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetTextStorageRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetTextStorageResponse{TextStorage: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetTextStorageRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetTextStorageResponse{TextStorage: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetTextStorageRequest{
				Index: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.TextStorage(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestTextStorageQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.ExampleKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNTextStorage(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllTextStorageRequest {
		return &types.QueryAllTextStorageRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.TextStorageAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.TextStorage), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.TextStorage),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.TextStorageAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.TextStorage), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.TextStorage),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.TextStorageAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.TextStorage),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.TextStorageAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
