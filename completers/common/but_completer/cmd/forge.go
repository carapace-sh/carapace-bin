package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var forgeCmd = &cobra.Command{
	Use:     "forge",
	Short:   "Commands for interacting with forges like GitHub, GitLab (coming soon), etc",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: "server interactions",
}

func init() {
	carapace.Gen(forgeCmd).Standalone()

	forgeCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(forgeCmd)
}
