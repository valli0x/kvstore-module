package kvstore

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "kvstore/api/kvstore/kvstore"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod:      "GetEntry",
					Use:            "get-entry [key]",
					Short:          "Query get-entry",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "key"}},
				},

				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "SetEntry",
					Use:            "set-entry [key] [value]",
					Short:          "Send a set-entry tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "key"}, {ProtoField: "value"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
