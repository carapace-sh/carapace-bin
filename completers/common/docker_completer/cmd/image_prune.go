package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var image_pruneCmd = &cobra.Command{
	Use:   "prune [OPTIONS]",
	Short: "Remove unused images",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_pruneCmd).Standalone()

	image_pruneCmd.Flags().BoolP("all", "a", false, "Remove all unused images, not just dangling ones")
	image_pruneCmd.Flags().String("filter", "", "Provide filter values (e.g. \"until=<timestamp>\")")
	image_pruneCmd.Flags().BoolP("force", "f", false, "Do not prompt for confirmation")
	imageCmd.AddCommand(image_pruneCmd)

	carapace.Gen(image_pruneCmd).FlagCompletion(carapace.ActionMap{
		"filter": carapace.ActionMultiPartsN("=", 2, func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return carapace.ActionValuesDescribed(
					"dangling", "Filter by dangling state",
					"label", "Filter by label (key or key=value)",
					"until", "Prune objects created before given duration",
				).Suffix("=")
			default:
				switch c.Parts[0] {
				case "dangling":
					return carapace.ActionValues("true", "false").StyleF(style.ForKeyword)
				default:
					return carapace.ActionValues()
				}
			}
		}),
	})
}
