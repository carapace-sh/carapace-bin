package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Aliases: []string{"r"},
	Short: "run environments",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runCmd).Standalone()

	add_common_flags(runCmd)
	add_common_subcommand_flags(runCmd)
	add_pkg_only_flags(runCmd)
	add_env_filtering_flags(runCmd)
	add_env_select_flag(runCmd)
	add_common_run_flags(runCmd)

	rootCmd.AddCommand(runCmd)
}
