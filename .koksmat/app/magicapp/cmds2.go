package magicapp

import (
	"github.com/spf13/cobra"

	"github.com/magicbutton/magic-devices/cmds"
	"github.com/magicbutton/magic-devices/utils"
)

func RegisterCmds() {
	utils.RootCmd.PersistentFlags().StringVarP(&utils.Output, "output", "o", "", "Output format (json, yaml, xml, etc.)")

	magicCmd := &cobra.Command{
		Use:   "magic",
		Short: "Magic",
		Long:  `Supporting the "App" domain`,
	}
	MagicKpiPostCmd := &cobra.Command{
		Use:   "kpi ",
		Short: "Process KPI",
		Long:  ``,
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()

			cmds.MagicKpiPost(ctx, args)
		},
	}
	magicCmd.AddCommand(MagicKpiPostCmd)

	utils.RootCmd.AddCommand(magicCmd)
	importCmd := &cobra.Command{
		Use:   "import",
		Short: "Import",
		Long:  `Supporting the "App" domain`,
	}
	ImportFullkpiPostCmd := &cobra.Command{
		Use:   "fullkpi ",
		Short: "Import Sheet FullKpi",
		Long:  ``,
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()

			cmds.ImportFullkpiPost(ctx, args)
		},
	}
	importCmd.AddCommand(ImportFullkpiPostCmd)
	ImportIntunePostCmd := &cobra.Command{
		Use:   "intune ",
		Short: "Import Sheet Intune",
		Long:  ``,
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()

			cmds.ImportIntunePost(ctx, args)
		},
	}
	importCmd.AddCommand(ImportIntunePostCmd)
	ImportTabellePostCmd := &cobra.Command{
		Use:   "tabelle ",
		Short: "Import Sheet Tabelle",
		Long:  ``,
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()

			cmds.ImportTabellePost(ctx, args)
		},
	}
	importCmd.AddCommand(ImportTabellePostCmd)

	utils.RootCmd.AddCommand(importCmd)
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
