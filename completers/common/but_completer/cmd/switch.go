package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var switchCmd = &cobra.Command{
	Use:    "switch",
	Short:  "Switch to a local branch, workspace branch ID, or the GitButler workspace",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(switchCmd).Standalone()

	switchCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	switchCmd.Flags().BoolP("new", "n", false, "Create a branch at the project target and switch to it")
	switchCmd.Flags().BoolP("workspace", "w", false, "Switch back to gitbutler/workspace")
	rootCmd.AddCommand(switchCmd)
}
