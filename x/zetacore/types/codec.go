package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	// this line is used by starport scaffolding # 2
	cdc.RegisterConcrete(&MsgCreateTSSVoter{}, "zetacore/CreateTSSVoter", nil)

	cdc.RegisterConcrete(&MsgGasBalanceVoter{}, "zetacore/GasBalanceVoter", nil)

	cdc.RegisterConcrete(&MsgGasPriceVoter{}, "zetacore/GasPriceVoter", nil)

	cdc.RegisterConcrete(&MsgNonceVoter{}, "zetacore/NonceVoter", nil)

	cdc.RegisterConcrete(&MsgReceiveConfirmation{}, "zetacore/ReceiveConfirmation", nil)

	cdc.RegisterConcrete(&MsgSendVoter{}, "zetacore/SendVoter", nil)

	cdc.RegisterConcrete(&MsgSetNodeKeys{}, "zetacore/SetNodeKeys", nil)

}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// this line is used by starport scaffolding # 3
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateTSSVoter{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgGasBalanceVoter{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgGasPriceVoter{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgNonceVoter{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgReceiveConfirmation{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSendVoter{},
	)

	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetNodeKeys{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	//amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)