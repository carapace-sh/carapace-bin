package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var file_dependencyCmd = &cobra.Command{
	Use:   "dependency",
	Short: "Manage file dependencies",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(file_dependencyCmd).Standalone()

	file_dependencyCmd.Flags().BoolP("help", "h", false, "Print help")
	fileCmd.AddCommand(file_dependencyCmd)
}
