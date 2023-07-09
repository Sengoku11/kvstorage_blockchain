package keeper

import (
	"context"

	"example/x/example/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateTextStorage(goCtx context.Context, msg *types.MsgCreateTextStorage) (*types.MsgCreateTextStorageResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetTextStorage(
		ctx,
		msg.Index,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var textStorage = types.TextStorage{
		Creator: msg.Creator,
		Index:   msg.Index,
		Text:    msg.Text,
	}

	k.SetTextStorage(
		ctx,
		textStorage,
	)
	return &types.MsgCreateTextStorageResponse{}, nil
}

func (k msgServer) UpdateTextStorage(goCtx context.Context, msg *types.MsgUpdateTextStorage) (*types.MsgUpdateTextStorageResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetTextStorage(
		ctx,
		msg.Index,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var textStorage = types.TextStorage{
		Creator: msg.Creator,
		Index:   msg.Index,
		Text:    msg.Text,
	}

	k.SetTextStorage(ctx, textStorage)

	return &types.MsgUpdateTextStorageResponse{}, nil
}

func (k msgServer) DeleteTextStorage(goCtx context.Context, msg *types.MsgDeleteTextStorage) (*types.MsgDeleteTextStorageResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetTextStorage(
		ctx,
		msg.Index,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveTextStorage(
		ctx,
		msg.Index,
	)

	return &types.MsgDeleteTextStorageResponse{}, nil
}
