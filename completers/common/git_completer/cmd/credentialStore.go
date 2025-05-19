package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var credentialStoreCmd = &cobra.Command{
	Use:     "credential-store",
	Short:   "Helper to store credentials on disk",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_helper].ID,
}

func init() {
	carapace.Gen(credentialStoreCmd).Standalone()

	credentialStoreCmd.Flags().String("file", "", "fetch and store credentials in <path>")
	rootCmd.AddCommand(credentialStoreCmd)

	carapace.Gen(credentialStoreCmd).FlagCompletion(carapace.ActionMap{
		"file": carapace.ActionFiles(),
	})
}
