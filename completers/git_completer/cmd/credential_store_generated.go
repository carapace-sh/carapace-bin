package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var credential_storeCmd = &cobra.Command{
	Use:     "credential-store",
	Short:   "Helper to store credentials on disk",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_helper].ID,
}

func init() {
	carapace.Gen(credential_storeCmd).Standalone()
	credential_storeCmd.Flags().String("file", "", "fetch and store credentials in <path>")
	rootCmd.AddCommand(credential_storeCmd)
}
