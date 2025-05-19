package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/yay"
	"github.com/carapace-sh/carapace-bin/pkg/util/embed"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "yay",
	Short: "An AUR Helper written in Go",
	Long:  "https://github.com/Jguer/yay",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()
	rootCmd.Flags().BoolP("version", "V", false, "show version")

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		yay.ActionPackageSearch(),
	)

	// TODO add missing flags (permanent configuration, "new" flags, ...)
	embed.SubcommandsAsFlags(rootCmd,
		buildCmd,
		databaseCmd,
		deptestCmd,
		filesCmd,
		getpkgbuildCmd,
		queryCmd,
		removeCmd,
		showCmd,
		syncCmd,
		upgradeCmd,
		webCmd,
		yayCmd,
	)
}
