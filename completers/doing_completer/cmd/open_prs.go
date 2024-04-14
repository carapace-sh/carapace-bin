package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var open_prsCmd = &cobra.Command{
	Use:   "prs",
	Short: "Open active PRs for repository view",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(open_prsCmd).Standalone()

	open_prsCmd.Flags().Bool("help", false, "Show this message and exit.")
	openCmd.AddCommand(open_prsCmd)
}
