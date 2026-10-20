package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bootstrap_filesCmd = &cobra.Command{
	Use:   "files",
	Short: "Manage system files and directories declared in `[bootstrap]`",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootstrap_filesCmd).Standalone()

	bootstrap_filesCmd.Flags().BoolP("help", "h", false, "Print help")
	bootstrapCmd.AddCommand(bootstrap_filesCmd)
}
