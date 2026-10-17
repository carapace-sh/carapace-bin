package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var configCheckCmd = &cobra.Command{
	Use:   "check [PATH]",
	Short: "Load the config file and report what it found",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configCheckCmd).Standalone()

	configCheckCmd.Flags().Bool("json", false, "output the result as JSON")

	configCmd.AddCommand(configCheckCmd)

	carapace.Gen(configCheckCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
