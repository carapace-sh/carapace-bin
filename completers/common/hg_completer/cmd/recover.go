package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var recoverCmd = &cobra.Command{
	Use:     "recover",
	Short:   "roll back an interrupted transaction",
	GroupID: groups[group_repository_maintenance].ID,
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(recoverCmd).Standalone()

	recoverCmd.Flags().Bool("verify", false, "run `hg verify` after successful recover")
	rootCmd.AddCommand(recoverCmd)
}
