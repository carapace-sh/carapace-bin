package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "test an archive",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(testCmd).Standalone()

	testCmd.Flags().String("password", "", "password for encrypted files")
	rootCmd.AddCommand(testCmd)

	carapace.Gen(testCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
