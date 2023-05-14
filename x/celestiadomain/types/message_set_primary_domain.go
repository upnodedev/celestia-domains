package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetPrimaryDomain = "set_primary_domain"

var _ sdk.Msg = &MsgSetPrimaryDomain{}

func NewMsgSetPrimaryDomain(creator string, domain string) *MsgSetPrimaryDomain {
	return &MsgSetPrimaryDomain{
		Creator: creator,
		Domain:  domain,
	}
}

func (msg *MsgSetPrimaryDomain) Route() string {
	return RouterKey
}

func (msg *MsgSetPrimaryDomain) Type() string {
	return TypeMsgSetPrimaryDomain
}

func (msg *MsgSetPrimaryDomain) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetPrimaryDomain) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetPrimaryDomain) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
