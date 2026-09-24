package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var selectCmd = &cobra.Command{
	Use:    "select",
	Short:  "Deprecated: use wt switch instead",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(selectCmd).Standalone()

	selectCmd.Flags().Bool("branches", false, "Include branches without worktrees")
	selectCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	selectCmd.Flags().Bool("remotes", false, "Include remote branches")
	rootCmd.AddCommand(selectCmd)

}
