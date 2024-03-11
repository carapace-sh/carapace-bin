package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gitwebCmd = &cobra.Command{
	Use:     "gitweb",
	Short:   "Git web interface (web frontend to Git repositories)",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_interrogator].ID,
}

func init() {
	carapace.Gen(gitwebCmd).Standalone()

	rootCmd.AddCommand(gitwebCmd)
}
