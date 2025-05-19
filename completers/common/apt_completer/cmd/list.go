package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list [pattern]...",
	Short: "list packages based on package names",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listCmd).Standalone()

	listCmd.Flags().BoolP("all-versions", "a", false, "print full records for all available versions")
	listCmd.Flags().Bool("installed", false, "list installed packages")
	listCmd.Flags().Bool("manual-installed", false, "list manually installed packages")
	listCmd.Flags().StringP("target-release", "t", "", "which distribution packages to retrieved from")
	listCmd.Flags().Bool("upgradable", false, "list upgradable packages")
	listCmd.Flags().BoolP("verbose", "v", false, "verbose output")
	rootCmd.AddCommand(listCmd)
}
