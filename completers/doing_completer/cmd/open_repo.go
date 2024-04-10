package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var open_repoCmd = &cobra.Command{
	Use:   "repo",
	Short: "Open repository view",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(open_repoCmd).Standalone()

	open_repoCmd.Flags().Bool("help", false, "Show this message and exit.")
	openCmd.AddCommand(open_repoCmd)
}
