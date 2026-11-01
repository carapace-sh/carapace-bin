package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var runtimeCmd = &cobra.Command{
	Use:     "runtime",
	Short:   "Manage runtimes",
	Aliases: []string{"rt"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runtimeCmd).Standalone()

	runtimeCmd.Flags().BoolP("global", "g", false, "Install the runtime globally")
	runtimeCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	runtimeCmd.Flags().BoolP("save-dev", "D", false, "Save the runtime to `devEngines.runtime`. This is the default")
	runtimeCmd.Flags().BoolP("save-prod", "P", false, "Save the runtime to `engines.runtime`")
	rootCmd.AddCommand(runtimeCmd)
}
