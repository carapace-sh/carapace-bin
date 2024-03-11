package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/yarn"
	"github.com/spf13/cobra"
)

var unplugCmd = &cobra.Command{
	Use:     "unplug",
	Short:   "Force the unpacking of a list of packages",
	GroupID: "general",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(unplugCmd).Standalone()

	unplugCmd.Flags().BoolP("all", "A", false, "Unplug direct dependencies from the entire project")
	unplugCmd.Flags().Bool("json", false, "Format the output as an NDJSON stream")
	unplugCmd.Flags().BoolP("recursive", "R", false, "Unplug both direct and transitive dependencies")
	rootCmd.AddCommand(unplugCmd)

	carapace.Gen(unplugCmd).PositionalAnyCompletion(
		yarn.ActionDependencies(),
	)
}
