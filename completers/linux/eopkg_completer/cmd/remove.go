package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:     "remove",
	Short:   "Remove packages",
	Aliases: []string{"rm"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(removeCmd).Standalone()

	removeCmd.Flags().Bool("ignore-dependency", false, "Do not take dependency info into account")
	removeCmd.Flags().Bool("ignore-comar", false, "Bypass COMAR configuration agent")
	removeCmd.Flags().Bool("ignore-safety", false, "Bypass safety switch")
	removeCmd.Flags().BoolP("dry-run", "n", false, "Do not perform any action, just show what would be done")
	removeCmd.Flags().Bool("purge", false, "Remove package and purge")
	removeCmd.Flags().StringP("component", "c", "", "Remove component's packages")

	rootCmd.AddCommand(removeCmd)
}
