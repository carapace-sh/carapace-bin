package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var scpCmd = &cobra.Command{
	Use:   "scp",
	Short: "copy files to or from a machine via SCP",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(scpCmd).Standalone()

	scpCmd.Flags().Bool("from", false, "Copy a file from specified machine")
	scpCmd.Flags().Bool("to", false, "Copy a file to specified machine")

	rootCmd.AddCommand(scpCmd)

	carapace.Gen(scpCmd).PositionalCompletion(
		carapace.ActionValues(), // machine
		carapace.ActionFiles(),
		carapace.ActionFiles(),
	)
}
