package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create an archive",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(createCmd).Standalone()

	createCmd.Flags().String("password", "", "password to encrypt files")
	rootCmd.AddCommand(createCmd)

	carapace.Gen(createCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
