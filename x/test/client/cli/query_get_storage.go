package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"test/x/test/types"
)

var _ = strconv.Itoa(0)

func CmdGetStorage() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-storage [address] [position] [block-tag]",
		Short: "Query getStorage",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqAddress := args[0]
			reqPosition := args[1]
			reqBlockTag := args[2]

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetStorageRequest{

				Address:  reqAddress,
				Position: reqPosition,
				BlockTag: reqBlockTag,
			}

			res, err := queryClient.GetStorage(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
