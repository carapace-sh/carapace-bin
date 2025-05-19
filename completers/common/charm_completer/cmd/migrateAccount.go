package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrateAccountCmd = &cobra.Command{
	Use:    "migrate-account",
	Short:  "",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrateAccountCmd).Standalone()

	migrateAccountCmd.Flags().BoolP("verbose", "v", false, "print debug output")
	rootCmd.AddCommand(migrateAccountCmd)
}
