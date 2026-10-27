package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_patchIdCmd = &cobra.Command{
	Use:   "patch-id",
	Short: "Print stable patch-IDs for one or more commits",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_patchIdCmd).Standalone()

	helpCmd.AddCommand(help_patchIdCmd)
}
