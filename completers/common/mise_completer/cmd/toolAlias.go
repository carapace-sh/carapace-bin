package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/mise"
	"github.com/spf13/cobra"
)

var toolAliasCmd = &cobra.Command{
	Use:     "tool-alias",
	Short:   "Manage tool backend and version aliases",
	Aliases: []string{"alias", "aliases"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(toolAliasCmd).Standalone()

	toolAliasCmd.Flags().BoolP("help", "h", false, "Print help")
	toolAliasCmd.Flags().Bool("no-header", false, "Do not print the table header")
	toolAliasCmd.Flags().String("plugin", "", "Filter aliases by tool")
	toolAliasCmd.Flags().StringP("tool", "p", "", "Filter aliases by tool")
	toolAliasCmd.Flag("plugin").Hidden = true
	rootCmd.AddCommand(toolAliasCmd)

	carapace.Gen(toolAliasCmd).FlagCompletion(carapace.ActionMap{
		"tool": mise.ActionTools(),
	})
}
