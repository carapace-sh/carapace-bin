package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/mise"
	"github.com/spf13/cobra"
)

var installs_migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Reinstall legacy installations into the identity install layout",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(installs_migrateCmd).Standalone()

	installs_migrateCmd.Flags().BoolP("dry-run", "n", false, "Show what would move without changing anything")
	installs_migrateCmd.Flags().BoolP("help", "h", false, "Print help")
	installsCmd.AddCommand(installs_migrateCmd)

	carapace.Gen(installs_migrateCmd).PositionalAnyCompletion(
		mise.ActionToolVersions(),
	)
}
