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
	downloadCmd := &cobra.Command{
		Use:   "download",
		Short: "Download",
		Long:  `Supporting the "App" domain`,
	}
	DownloadExcelPostCmd := &cobra.Command{
		Use:   "excel ",
		Short: "Download excel",
		Long:  ``,
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()

			cmds.DownloadExcelPost(ctx, args)
		},
	}
	downloadCmd.AddCommand(DownloadExcelPostCmd)

	utils.RootCmd.AddCommand(downloadCmd)
	extractCmd := &cobra.Command{
		Use:   "extract",
		Short: "Extract",
		Long:  `Supporting the "App" domain`,
	}
	ExtractFullkpiPostCmd := &cobra.Command{
		Use:   "fullkpi ",
		Short: "Import Sheet FullKpi",
		Long:  ``,
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()

			cmds.ExtractFullkpiPost(ctx, args)
		},
	}
	extractCmd.AddCommand(ExtractFullkpiPostCmd)
	ExtractIntunePostCmd := &cobra.Command{
		Use:   "intune ",
		Short: "Import Sheet Intune",
		Long:  ``,
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()

			cmds.ExtractIntunePost(ctx, args)
		},
	}
	extractCmd.AddCommand(ExtractIntunePostCmd)
	ExtractTabellePostCmd := &cobra.Command{
		Use:   "tabelle ",
		Short: "Import Sheet Tabelle",
		Long:  ``,
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()

			cmds.ExtractTabellePost(ctx, args)
		},
	}
	extractCmd.AddCommand(ExtractTabellePostCmd)
	ExtractDownloadExcelPostCmd := &cobra.Command{
		Use:   "download-excel ",
		Short: "Download Excel",
		Long:  ``,
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()

			cmds.ExtractDownloadExcelPost(ctx, args)
		},
	}
	extractCmd.AddCommand(ExtractDownloadExcelPostCmd)

	utils.RootCmd.AddCommand(extractCmd)
	insertCmd := &cobra.Command{
		Use:   "insert",
		Short: "Insert",
		Long:  `Supporting the "App" domain`,
	}
	InsertSqlPostCmd := &cobra.Command{
		Use:   "sql ",
		Short: "Insert SQL",
		Long:  ``,
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()

			cmds.InsertSqlPost(ctx, args)
		},
	}
	insertCmd.AddCommand(InsertSqlPostCmd)

	utils.RootCmd.AddCommand(insertCmd)
	transformCmd := &cobra.Command{
		Use:   "transform",
		Short: "Transform",
		Long:  `Supporting the "App" domain`,
	}
	TransformDownloadExcelPostCmd := &cobra.Command{
		Use:   "download-excel ",
		Short: "Download Excel",
		Long:  ``,
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()

			cmds.TransformDownloadExcelPost(ctx, args)
		},
	}
	transformCmd.AddCommand(TransformDownloadExcelPostCmd)

	utils.RootCmd.AddCommand(transformCmd)
	loadCmd := &cobra.Command{
		Use:   "load",
		Short: "Load",
		Long:  `Supporting the "App" domain`,
	}
	LoadLoadPostCmd := &cobra.Command{
		Use:   "load ",
		Short: "Load",
		Long:  ``,
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()

			cmds.LoadLoadPost(ctx, args)
		},
	}
	loadCmd.AddCommand(LoadLoadPostCmd)

	utils.RootCmd.AddCommand(loadCmd)
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
