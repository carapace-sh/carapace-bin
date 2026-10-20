package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/mise"
	"github.com/spf13/cobra"
)

var lsCmd = &cobra.Command{
	Use:     "ls",
	Short:   "List installed and active tool versions",
	Aliases: []string{"list"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lsCmd).Standalone()

	lsCmd.Flags().Bool("all-sources", false, "Show requests from every config file mise has seen, not only loaded ones")
	lsCmd.Flags().StringSliceP("backend", "b", nil, "Only show tools from this backend, such as aqua, cargo, core, or go")
	lsCmd.Flags().BoolP("current", "c", false, "Only show versions the loaded config files request")
	lsCmd.Flags().BoolP("global", "g", false, "Only show versions the global config requests")
	lsCmd.Flags().Bool("grouped", false, "List tools in a separate section for each backend")
	lsCmd.Flags().BoolP("help", "h", false, "Print help")
	lsCmd.Flags().BoolP("installed", "i", false, "Only show installed versions")
	lsCmd.Flags().BoolP("json", "J", false, "Output in JSON format")
	lsCmd.Flags().BoolP("local", "l", false, "Only show versions the local config requests")
	lsCmd.Flags().BoolP("missing", "m", false, "Only show requested versions that are not installed")
	lsCmd.Flags().Bool("monorepo", false, "Also list tools from every config root in `[monorepo].config_roots`")
	lsCmd.Flags().Bool("no-header", false, "Do not print the header row")
	lsCmd.Flags().Bool("no-headers", false, "Do not print the header row")
	lsCmd.Flags().Bool("no-truncate", false, "Truncate long terminal output to fit the available width")
	lsCmd.Flags().BoolP("offline", "o", false, "Don't fetch information such as outdated versions")
	lsCmd.Flags().Bool("outdated", false, "Show whether each version is outdated")
	lsCmd.Flags().StringP("plugin", "p", "", "")
	lsCmd.Flags().String("prefix", "", "Only show versions starting with PREFIX (requires a tool argument)")
	lsCmd.Flags().Bool("prunable", false, "List only tools that can be pruned with `mise prune`")
	lsCmd.Flags().Bool("truncate", true, "Truncate long terminal output to fit the available width")
	lsCmd.Flag("no-headers").Hidden = true
	lsCmd.Flag("no-truncate").Hidden = true
	lsCmd.Flag("offline").Hidden = true
	lsCmd.Flag("plugin").Hidden = true
	rootCmd.AddCommand(lsCmd)

	carapace.Gen(lsCmd).FlagCompletion(carapace.ActionMap{
		"backend": mise.ActionBackends(),
		"plugin":  mise.ActionPlugins(),
	})

	carapace.Gen(lsCmd).PositionalAnyCompletion(
		mise.ActionInstalledTools(),
	)
}
