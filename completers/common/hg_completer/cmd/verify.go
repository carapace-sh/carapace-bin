package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var verifyCmd = &cobra.Command{
	Use:     "verify",
	Short:   "verify the integrity of the repository",
	GroupID: groups[group_repository_maintenance].ID,
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(verifyCmd).Standalone()

	verifyCmd.Flags().Bool("full", false, "perform more checks (EXPERIMENTAL)")
	rootCmd.AddCommand(verifyCmd)
}
