package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "test if a file or folder has a custom icon",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(testCmd).Standalone()

	testCmd.Flags().BoolP("quiet", "q", false, "silence status output")
	rootCmd.AddCommand(testCmd)

	carapace.Gen(testCmd).PositionalCompletion(carapace.ActionFiles())
}
