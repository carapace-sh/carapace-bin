package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var moduleCmd = &cobra.Command{
	Use:   "module <subcommand> [options] [<module-spec>...]",
	Short: "manage modules (deprecated)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(moduleCmd).Standalone()

	rootCmd.AddCommand(moduleCmd)
}
