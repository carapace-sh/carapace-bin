package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-pnpm/pkg/actions/tools/pnpm"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Runs a defined package script",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runCmd).Standalone()

	runCmd.Flags().Bool("dry-run", false, "Print the task graph a recursive run would execute, without running anything. Only meaningful together with the global `-r` / `--recursive` flag")
	runCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	runCmd.Flags().Bool("if-present", false, "Avoid exiting with a non-zero exit code when the script is undefined")
	runCmd.Flags().Bool("json", false, "With `--dry-run`, print the tasks and their resolved dependency edges as JSON")
	runCmd.Flags().BoolP("sequential", "s", false, "Run the specified scripts one by one")
	rootCmd.AddCommand(runCmd)

	carapace.Gen(runCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			filter, _ := runCmd.Flags().GetString("filter")
			return pnpm.ActionScriptsForFilter(filter)
		}),
	)
}
