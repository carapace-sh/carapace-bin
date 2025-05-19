package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pr_checkoutCmd = &cobra.Command{
	Use:   "checkout [OPTIONS] PR_ID",
	Short: "Check out a pull request in git",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pr_checkoutCmd).Standalone()

	pr_checkoutCmd.Flags().Bool("help", false, "Show this message and exit.")
	prCmd.AddCommand(pr_checkoutCmd)
}
