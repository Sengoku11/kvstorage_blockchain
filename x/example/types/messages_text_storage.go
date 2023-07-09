package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateTextStorage = "create_text_storage"
	TypeMsgUpdateTextStorage = "update_text_storage"
	TypeMsgDeleteTextStorage = "delete_text_storage"
)

var _ sdk.Msg = &MsgCreateTextStorage{}

func NewMsgCreateTextStorage(
	creator string,
	index string,
	text string,

) *MsgCreateTextStorage {
	return &MsgCreateTextStorage{
		Creator: creator,
		Index:   index,
		Text:    text,
	}
}

func (msg *MsgCreateTextStorage) Route() string {
	return RouterKey
}

func (msg *MsgCreateTextStorage) Type() string {
	return TypeMsgCreateTextStorage
}

func (msg *MsgCreateTextStorage) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateTextStorage) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateTextStorage) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateTextStorage{}

func NewMsgUpdateTextStorage(
	creator string,
	index string,
	text string,

) *MsgUpdateTextStorage {
	return &MsgUpdateTextStorage{
		Creator: creator,
		Index:   index,
		Text:    text,
	}
}

func (msg *MsgUpdateTextStorage) Route() string {
	return RouterKey
}

func (msg *MsgUpdateTextStorage) Type() string {
	return TypeMsgUpdateTextStorage
}

func (msg *MsgUpdateTextStorage) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateTextStorage) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateTextStorage) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteTextStorage{}

func NewMsgDeleteTextStorage(
	creator string,
	index string,

) *MsgDeleteTextStorage {
	return &MsgDeleteTextStorage{
		Creator: creator,
		Index:   index,
	}
}
func (msg *MsgDeleteTextStorage) Route() string {
	return RouterKey
}

func (msg *MsgDeleteTextStorage) Type() string {
	return TypeMsgDeleteTextStorage
}

func (msg *MsgDeleteTextStorage) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteTextStorage) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteTextStorage) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
