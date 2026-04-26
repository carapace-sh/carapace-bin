package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var aliasCmd = &cobra.Command{
	Use:   "alias [name]",
	Short: "Create a shell alias for gh stack",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(aliasCmd).Standalone()

	aliasCmd.Flags().Bool("remove", false, "Remove a previously created alias")
	rootCmd.AddCommand(aliasCmd)
}
