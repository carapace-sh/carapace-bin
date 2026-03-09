package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var versionlockCmd = &cobra.Command{
	Use:   "versionlock <subcommand> <package-spec>...",
	Short: "protect packages from updates to newer versions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(versionlockCmd).Standalone()

	rootCmd.AddCommand(versionlockCmd)
}
