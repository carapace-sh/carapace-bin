package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var completion_powershellCmd = &cobra.Command{
	Use:   "powershell",
	Short: "powershell completion.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(completion_powershellCmd).Standalone()

	completionCmd.AddCommand(completion_powershellCmd)
}
