package cli

import (
	"strconv"

	"test/x/test/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdGetBlockHash() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-block-hash [block-number]",
		Short: "Query getBlockHash",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqBlockNumber := args[0]

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			reqNumber, err := strconv.ParseUint(reqBlockNumber, 10, 64)
			if err != nil {
				return
			}
			params := &types.QueryGetBlockHashRequest{

				BlockNumber: reqNumber,
			}

			res, err := queryClient.GetBlockHash(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
