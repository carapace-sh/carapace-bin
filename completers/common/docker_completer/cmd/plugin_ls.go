package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var plugin_lsCmd = &cobra.Command{
	Use:     "ls [OPTIONS]",
	Short:   "List plugins",
	Aliases: []string{"list"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugin_lsCmd).Standalone()

	plugin_lsCmd.Flags().StringP("filter", "f", "", "Provide filter values (e.g. \"enabled=true\")")
	plugin_lsCmd.Flags().String("format", "", "Format output using a custom template:")
	plugin_lsCmd.Flags().Bool("no-trunc", false, "Don't truncate output")
	plugin_lsCmd.Flags().BoolP("quiet", "q", false, "Only display plugin IDs")
	pluginCmd.AddCommand(plugin_lsCmd)

	carapace.Gen(plugin_lsCmd).FlagCompletion(carapace.ActionMap{
		"filter": carapace.ActionMultiPartsN("=", 2, func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return carapace.ActionValuesDescribed(
					"capability", "Filter by capability",
					"enabled", "Filter by enabled state",
				).Suffix("=")
			default:
				switch c.Parts[0] {
				case "enabled":
					return carapace.ActionValues("true", "false").StyleF(style.ForKeyword)
				default:
					return carapace.ActionValues()
				}
			}
		}),
	})
}
