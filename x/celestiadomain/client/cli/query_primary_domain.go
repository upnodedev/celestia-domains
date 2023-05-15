package cli

import (
	"strconv"

	"celestia-domain/x/celestiadomain/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdPrimaryDomain() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "primary-domain [owner]",
		Short: "Query primaryDomain",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argOwner := args[0]

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryPrimaryDomainRequest{
				Owner: argOwner,
			}

			res, err := queryClient.PrimaryDomain(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
