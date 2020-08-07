package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/ovrclk/akash/x/deployment/types"
	"github.com/spf13/cobra"
)

// GetQueryCmd returns the query commands for the deployment module
func GetQueryCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Deployment query commands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		cmdDeployments(),
		cmdDeployment(),
		getGroupCmd(),
	)

	return cmd
}

func cmdDeployments() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Query for all deployments",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadQueryCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			dfilters, state, err := DepFiltersFromFlags(cmd.Flags())
			if err != nil {
				return err
			}

			// checking state flag
			stateVal, ok := types.Deployment_State_value[state]

			if (!ok && (state != "")) || state == "invalid" {
				return ErrStateValue
			}

			dfilters.State = types.Deployment_State(stateVal)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			params := &types.QueryDeploymentsRequest{
				Filters:    dfilters,
				Pagination: pageReq,
			}

			res, err := queryClient.Deployments(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintOutput(res.Deployments)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "deployments")
	AddDeploymentFilterFlags(cmd.Flags())

	return cmd
}

func cmdDeployment() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Query deployment",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadQueryCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			id, err := DeploymentIDFromFlags(cmd.Flags(), "")
			if err != nil {
				return err
			}

			res, err := queryClient.Deployment(context.Background(), &types.QueryDeploymentRequest{ID: id})
			if err != nil {
				return err
			}

			return clientCtx.PrintOutput(res.Deployment)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	AddDeploymentIDFlags(cmd.Flags())
	MarkReqDeploymentIDFlags(cmd)

	return cmd
}

func getGroupCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "group",
		Short:                      "Deployment group query commands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		cmdGetGroup(),
	)

	return cmd
}

func cmdGetGroup() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Query group of deployment",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadQueryCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			id, err := GroupIDFromFlags(cmd.Flags())
			if err != nil {
				return err
			}

			res, err := queryClient.Group(context.Background(), &types.QueryGroupRequest{ID: id})
			if err != nil {
				return err
			}

			return clientCtx.PrintOutput(res.Group)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	AddGroupIDFlags(cmd.Flags())
	MarkReqGroupIDFlags(cmd)

	return cmd
}
