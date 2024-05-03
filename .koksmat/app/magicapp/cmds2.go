package magicapp

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/magicbutton/magic-devices/cmds"
	"github.com/magicbutton/magic-devices/utils"
)

func RegisterCmds() {
	utils.RootCmd.PersistentFlags().StringVarP(&utils.Output, "output", "o", "", "Output format (json, yaml, xml, etc.)")

	healthCmd := &cobra.Command{
		Use:   "health",
		Short: "Health",
		Long:  `Describe the main purpose of this kitchen`,
	}
	HealthPingPostCmd := &cobra.Command{
		Use:   "ping  pong",
		Short: "Ping",
		Long:  `Simple ping endpoint`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()

			cmds.HealthPingPost(ctx, args)
		},
	}
	healthCmd.AddCommand(HealthPingPostCmd)
	HealthCoreversionPostCmd := &cobra.Command{
		Use:   "coreversion ",
		Short: "Core Version",
		Long:  ``,
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()

			cmds.HealthCoreversionPost(ctx, args)
		},
	}
	healthCmd.AddCommand(HealthCoreversionPostCmd)

	utils.RootCmd.AddCommand(healthCmd)
	usecasesCmd := &cobra.Command{
		Use:   "usecases",
		Short: "Use Cases",
		Long:  `Describe the main purpose of this kitchen`,
	}
	UsecasesGetGroupsPostCmd := &cobra.Command{
		Use:   "get-groups ",
		Short: "Get Groups",
		Long:  ``,
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()

			cmds.UsecasesGetGroupsPost(ctx, args)
		},
	}
	usecasesCmd.AddCommand(UsecasesGetGroupsPostCmd)
	UsecasesGetSegmentsPostCmd := &cobra.Command{
		Use:   "get-segments ",
		Short: "Get Segments",
		Long:  ``,
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()

			cmds.UsecasesGetSegmentsPost(ctx, args)
		},
	}
	usecasesCmd.AddCommand(UsecasesGetSegmentsPostCmd)
	UsecasesGetUserPostCmd := &cobra.Command{
		Use:   "get-user  userid",
		Short: "Get  User",
		Long:  ``,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()

			cmds.UsecasesGetUserPost(ctx, args)
		},
	}
	usecasesCmd.AddCommand(UsecasesGetUserPostCmd)

	utils.RootCmd.AddCommand(usecasesCmd)
	analyseCmd := &cobra.Command{
		Use:   "analyse",
		Short: "30-analyse",
		Long:  `Describe the main purpose of this kitchen`,
	}
	AnalyseParseGroupsPostCmd := &cobra.Command{
		Use:   "parse-groups ",
		Short: "Parse Groups",
		Long:  ``,
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()
			body, err := os.ReadFile(args[0])
			if err != nil {
				panic(err)
			}

			cmds.AnalyseParseGroupsPost(ctx, body, args)
		},
	}
	analyseCmd.AddCommand(AnalyseParseGroupsPostCmd)
	AnalyseParseUsersPostCmd := &cobra.Command{
		Use:   "parse-users ",
		Short: "Parse Users",
		Long:  ``,
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()
			body, err := os.ReadFile(args[0])
			if err != nil {
				panic(err)
			}

			cmds.AnalyseParseUsersPost(ctx, body, args)
		},
	}
	analyseCmd.AddCommand(AnalyseParseUsersPostCmd)

	utils.RootCmd.AddCommand(analyseCmd)
	provisionCmd := &cobra.Command{
		Use:   "provision",
		Short: "Provision",
		Long:  `Describe the main purpose of this kitchen`,
	}
	ProvisionAppdeployproductionPostCmd := &cobra.Command{
		Use:   "appdeployproduction ",
		Short: "App deploy to production",
		Long:  ``,
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()

			cmds.ProvisionAppdeployproductionPost(ctx, args)
		},
	}
	provisionCmd.AddCommand(ProvisionAppdeployproductionPostCmd)
	ProvisionWebdeployproductionPostCmd := &cobra.Command{
		Use:   "webdeployproduction ",
		Short: "Web deploy to production",
		Long:  ``,
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()

			cmds.ProvisionWebdeployproductionPost(ctx, args)
		},
	}
	provisionCmd.AddCommand(ProvisionWebdeployproductionPostCmd)
	ProvisionWebdeploytestPostCmd := &cobra.Command{
		Use:   "webdeploytest ",
		Short: "Web deploy to Test",
		Long:  ``,
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()

			cmds.ProvisionWebdeploytestPost(ctx, args)
		},
	}
	provisionCmd.AddCommand(ProvisionWebdeploytestPostCmd)

	utils.RootCmd.AddCommand(provisionCmd)
}
