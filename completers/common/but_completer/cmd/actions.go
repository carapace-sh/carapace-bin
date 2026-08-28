package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var actionsCmd = &cobra.Command{
	Use:    "actions",
	Short:  "INTERNAL: GitButler Actions are automated tasks (like macros) that can be performed on a repository",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(actionsCmd).Standalone()

	actionsCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(actionsCmd)
}