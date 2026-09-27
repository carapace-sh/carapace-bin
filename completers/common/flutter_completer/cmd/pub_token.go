package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pub_tokenCmd = &cobra.Command{
	Use:   "token",
	Short: "Manage authentication tokens for package repositories",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pub_tokenCmd).Standalone()

	pub_tokenCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	pubCmd.AddCommand(pub_tokenCmd)
}
