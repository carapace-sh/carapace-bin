package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var account_linkCmd = &cobra.Command{
	Use:   "link",
	Short: "Link your CLI sync account to your Hub account",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(account_linkCmd).Standalone()

	account_linkCmd.Flags().BoolP("help", "h", false, "Print help")
	accountCmd.AddCommand(account_linkCmd)
}
