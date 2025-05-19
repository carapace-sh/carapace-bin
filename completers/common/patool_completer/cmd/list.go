package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list members or one or more archives",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listCmd).Standalone()

	listCmd.Flags().String("password", "", "password for encrypted files")
	rootCmd.AddCommand(listCmd)

	carapace.Gen(listCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
