package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pub_token_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List servers for which a token exists",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pub_token_listCmd).Standalone()

	pub_token_listCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	pub_tokenCmd.AddCommand(pub_token_listCmd)
}
