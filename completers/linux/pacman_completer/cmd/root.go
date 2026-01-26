package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/util/embed"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pacman",
	Short: "package manager utility",
	Long:  "https://wiki.archlinux.de/title/Pacman",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("help", "h", false, "Display syntax for the given operation")
	rootCmd.Flags().BoolP("version", "V", false, "Display version")

	embed.SubcommandsAsFlags(rootCmd,
		databaseCmd,
		deptestCmd,
		filesCmd,
		queryCmd,
		removeCmd,
		syncCmd,
		upgradeCmd,
	)
}
