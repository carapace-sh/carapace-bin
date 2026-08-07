package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var orbitCmd = &cobra.Command{
	Use:   "orbit <command> [flags]",
	Short: "GitLab Knowledge Graph commands. (EXPERIMENTAL)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(orbitCmd).Standalone()

	rootCmd.AddCommand(orbitCmd)
}
