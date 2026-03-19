package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_account_linkCmd = &cobra.Command{
	Use:   "link",
	Short: "Link your CLI sync account to your Hub account",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_account_linkCmd).Standalone()

	help_accountCmd.AddCommand(help_account_linkCmd)
}
