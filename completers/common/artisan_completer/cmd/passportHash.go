package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var passportHashCmd = &cobra.Command{
	Use:   "passport:hash [--force]",
	Short: "Hash all of the existing secrets in the clients table",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(passportHashCmd).Standalone()

	passportHashCmd.Flags().Bool("force", false, "Force the operation to run without confirmation prompt")
	rootCmd.AddCommand(passportHashCmd)
}
