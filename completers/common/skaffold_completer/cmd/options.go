package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var optionsCmd = &cobra.Command{
	Use:   "options",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(optionsCmd).Standalone()

	applyCommonFlagCompletion(optionsCmd)
	rootCmd.AddCommand(optionsCmd)
}
