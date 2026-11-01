package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pub_token_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove secret token for package repository at <hosted-url>",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pub_token_removeCmd).Standalone()

	pub_token_removeCmd.Flags().Bool("all", false, "Remove all secret tokens.")
	pub_token_removeCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	pub_tokenCmd.AddCommand(pub_token_removeCmd)
}
