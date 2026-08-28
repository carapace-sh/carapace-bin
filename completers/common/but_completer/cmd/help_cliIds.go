package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_cliIdsCmd = &cobra.Command{
	Use:   "cli-ids",
	Short: "Smart IDs to reference commits, branches and more in but",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_cliIdsCmd).Standalone()

	help_cliIdsCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	helpCmd.AddCommand(help_cliIdsCmd)
}