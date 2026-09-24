package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debugupgraderepoCmd = &cobra.Command{
	Use:    "debugupgraderepo",
	Short:  "upgrade a repository to use different features",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debugupgraderepoCmd).Standalone()

	debugupgraderepoCmd.Flags().Bool("backup", false, "keep the old repository content around")
	debugupgraderepoCmd.Flags().Bool("changelog", false, "select the changelog for upgrade")
	debugupgraderepoCmd.Flags().Bool("filelogs", false, "select all filelogs for upgrade")
	debugupgraderepoCmd.Flags().Bool("manifest", false, "select the manifest for upgrade")
	debugupgraderepoCmd.Flags().StringArrayP("optimize", "o", nil, "extra optimization to perform")
	debugupgraderepoCmd.Flags().Bool("run", false, "performs an upgrade")
	rootCmd.AddCommand(debugupgraderepoCmd)
}
