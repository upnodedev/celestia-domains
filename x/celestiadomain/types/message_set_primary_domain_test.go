package types

import (
	"testing"

	"celestia-domain/testutil/sample"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgSetPrimaryDomain_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgSetPrimaryDomain
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgSetPrimaryDomain{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgSetPrimaryDomain{
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
