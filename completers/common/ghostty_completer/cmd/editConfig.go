package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var editConfigCmd = &cobra.Command{
	Use:   "+edit-config",
	Short: "The `edit-config` command opens the Ghostty configuration file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(editConfigCmd).Standalone()

	rootCmd.AddCommand(editConfigCmd)
}
