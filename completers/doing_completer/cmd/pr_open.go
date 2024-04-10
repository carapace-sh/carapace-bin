package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pr_openCmd = &cobra.Command{
	Use:   "open [OPTIONS] [PULLREQUEST_ID]",
	Short: "Alias: `doing open pr`",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pr_openCmd).Standalone()

	pr_openCmd.Flags().Bool("help", false, "Show this message and exit.")
	prCmd.AddCommand(pr_openCmd)
}
