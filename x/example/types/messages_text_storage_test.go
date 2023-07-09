package types

import (
	"testing"

	"example/testutil/sample"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateTextStorage_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateTextStorage
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateTextStorage{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateTextStorage{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgUpdateTextStorage_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateTextStorage
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateTextStorage{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateTextStorage{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgDeleteTextStorage_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteTextStorage
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteTextStorage{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteTextStorage{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
