package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var autoremoveCmd = &cobra.Command{
	Use:     "autoremove",
	Short:   "Remove packages and dependencies",
	Aliases: []string{"rmf"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoremoveCmd).Standalone()

	autoremoveCmd.Flags().Bool("ignore-dependency", false, "Do not take dependency info into account")
	autoremoveCmd.Flags().Bool("ignore-comar", false, "Bypass COMAR configuration agent")
	autoremoveCmd.Flags().Bool("ignore-safety", false, "Bypass safety switch")
	autoremoveCmd.Flags().BoolP("dry-run", "n", false, "Do not perform any action, just show what would be done")
	autoremoveCmd.Flags().Bool("purge", false, "Remove package and purge")

	rootCmd.AddCommand(autoremoveCmd)
}
