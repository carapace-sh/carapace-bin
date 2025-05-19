package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var open_prCmd = &cobra.Command{
	Use:   "pr",
	Short: "Open a specific PULLREQUEST_ID",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(open_prCmd).Standalone()

	open_prCmd.Flags().Bool("help", false, "Show this message and exit.")
	openCmd.AddCommand(open_prCmd)
}
