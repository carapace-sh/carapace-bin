package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var system_pruneCmd = &cobra.Command{
	Use:   "prune [OPTIONS]",
	Short: "Remove unused data",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(system_pruneCmd).Standalone()

	system_pruneCmd.Flags().BoolP("all", "a", false, "Remove all unused images not just dangling ones")
	system_pruneCmd.Flags().String("filter", "", "Provide filter values (e.g. \"label=<key>=<value>\")")
	system_pruneCmd.Flags().BoolP("force", "f", false, "Do not prompt for confirmation")
	system_pruneCmd.Flags().Bool("volumes", false, "Prune anonymous volumes")
	systemCmd.AddCommand(system_pruneCmd)

	carapace.Gen(system_pruneCmd).FlagCompletion(carapace.ActionMap{
		"filter": carapace.ActionMultiPartsN("=", 2, func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return carapace.ActionValuesDescribed(
					"label", "Filter by label (key or key=value)",
					"until", "Prune objects created before given duration",
				).Suffix("=")
			default:
				return carapace.ActionValues()
			}
		}),
	})
}
