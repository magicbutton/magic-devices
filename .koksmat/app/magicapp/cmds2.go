package magicapp

import (
	"github.com/spf13/cobra"

	"github.com/magicbutton/magic-devices/cmds"
	"github.com/magicbutton/magic-devices/utils"
)

func RegisterCmds() {
	utils.RootCmd.PersistentFlags().StringVarP(&utils.Output, "output", "o", "", "Output format (json, yaml, xml, etc.)")

	runCmd := &cobra.Command{
		Use:   "run",
		Short: "Processs",
		Long:  `Supporting the "App" domain`,
	}
	RunDevicekpiPostCmd := &cobra.Command{
		Use:   "devicekpi ",
		Short: "Get KPI report data",
		Long:  ``,
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()

			cmds.RunDevicekpiPost(ctx, args)
		},
	}
	runCmd.AddCommand(RunDevicekpiPostCmd)

	utils.RootCmd.AddCommand(runCmd)
	provisionCmd := &cobra.Command{
		Use:   "provision",
		Short: "Provision",
		Long:  `Supporting the "App" domain`,
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
	ProvisionJobdeployproductionPostCmd := &cobra.Command{
		Use:   "jobdeployproduction ",
		Short: "Job deploy to production",
		Long:  ``,
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()

			cmds.ProvisionJobdeployproductionPost(ctx, args)
		},
	}
	provisionCmd.AddCommand(ProvisionJobdeployproductionPostCmd)

	utils.RootCmd.AddCommand(provisionCmd)
	healthCmd := &cobra.Command{
		Use:   "health",
		Short: "Health",
		Long:  `Supporting the "App" domain`,
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
}
