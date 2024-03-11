package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var image_lsCmd = &cobra.Command{
	Use:     "ls [OPTIONS] [REPOSITORY[:TAG]]",
	Short:   "List images",
	Aliases: []string{"list"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_lsCmd).Standalone()

	image_lsCmd.Flags().BoolP("all", "a", false, "Show all images (default hides intermediate images)")
	image_lsCmd.Flags().Bool("digests", false, "Show digests")
	image_lsCmd.Flags().StringP("filter", "f", "", "Filter output based on conditions provided")
	image_lsCmd.Flags().String("format", "", "Format output using a custom template:")
	image_lsCmd.Flags().Bool("no-trunc", false, "Don't truncate output")
	image_lsCmd.Flags().BoolP("quiet", "q", false, "Only show image IDs")
	imageCmd.AddCommand(image_lsCmd)

	carapace.Gen(image_lsCmd).FlagCompletion(carapace.ActionMap{
		"filter": carapace.ActionMultiParts("=", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return carapace.ActionValuesDescribed(
					"dangling", "filter images by dangling state",
					"label", "filter images by label",
					"before", "filter images created before given id or references",
					"since", "filter images created since given id or references",
					"reference", "filter images whose reference matches the specified pattern",
				).Invoke(c).Suffix("=").ToA()
			case 1:
				switch c.Parts[0] {
				case "dangling":
					return carapace.ActionValues("true", "false")
				case "before":
					return docker.ActionRepositoryTags()
				case "since":
					return docker.ActionRepositoryTags()
				default:
					return carapace.ActionValues()
				}
			default:
				return carapace.ActionValues()
			}
		}),
	})

	carapace.Gen(image_lsCmd).PositionalCompletion(
		docker.ActionRepositoryTags(),
	)
}
