package keeper

import (
	"context"

	"example/x/example/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) TextStorageAll(goCtx context.Context, req *types.QueryAllTextStorageRequest) (*types.QueryAllTextStorageResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var textStorages []types.TextStorage
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	textStorageStore := prefix.NewStore(store, types.KeyPrefix(types.TextStorageKeyPrefix))

	pageRes, err := query.Paginate(textStorageStore, req.Pagination, func(key []byte, value []byte) error {
		var textStorage types.TextStorage
		if err := k.cdc.Unmarshal(value, &textStorage); err != nil {
			return err
		}

		textStorages = append(textStorages, textStorage)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllTextStorageResponse{TextStorage: textStorages, Pagination: pageRes}, nil
}

func (k Keeper) TextStorage(goCtx context.Context, req *types.QueryGetTextStorageRequest) (*types.QueryGetTextStorageResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetTextStorage(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetTextStorageResponse{TextStorage: val}, nil
}
