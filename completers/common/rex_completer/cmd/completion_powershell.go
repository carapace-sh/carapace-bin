package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var completionPowershellCmd = &cobra.Command{
	Use:   "powershell",
	Short: "Generate PowerShell completion",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(completionPowershellCmd).Standalone()

	completionCmd.AddCommand(completionPowershellCmd)
}
