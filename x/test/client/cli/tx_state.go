package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"test/x/test/types"
)

var _ = strconv.Itoa(0)

func CmdState() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "state [address] [position] [blocktag] [storage]",
		Short: "Broadcast message state",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argAddress := args[0]
			argPosition := args[1]
			argBlocktag := args[2]
			argStorage := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgState(
				clientCtx.GetFromAddress().String(),
				argAddress,
				argPosition,
				argBlocktag,
				argStorage,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
