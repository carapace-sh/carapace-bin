package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var imagesCmd = &cobra.Command{
	Use:     "images [OPTIONS] [REPOSITORY[:TAG]]",
	Short:   "List images",
	GroupID: "common",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imagesCmd).Standalone()

	imagesCmd.Flags().BoolP("all", "a", false, "Show all images (default hides intermediate and dangling images)")
	imagesCmd.Flags().Bool("digests", false, "Show digests")
	imagesCmd.Flags().StringP("filter", "f", "", "Filter output based on conditions provided")
	imagesCmd.Flags().String("format", "", "Format output using a custom template:")
	imagesCmd.Flags().Bool("no-trunc", false, "Don't truncate output")
	imagesCmd.Flags().BoolP("quiet", "q", false, "Only show image IDs")
	imagesCmd.Flags().Bool("tree", false, "List multi-platform images as a tree (EXPERIMENTAL)")
	rootCmd.AddCommand(imagesCmd)

	carapace.Gen(imagesCmd).FlagCompletion(carapace.ActionMap{
		"filter": carapace.ActionMultiPartsN("=", 2, func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return carapace.ActionValuesDescribed(
					"dangling", "filter images by dangling state",
					"label", "filter images by label",
					"before", "filter images created before given id or references",
					"since", "filter images created since given id or references",
					"reference", "filter images whose reference matches the specified pattern",
				).Suffix("=")
			default:
				switch c.Parts[0] {
				case "dangling":
					return carapace.ActionValues("true", "false").StyleF(style.ForKeyword)
				case "before":
					return docker.ActionRepositoryTags()
				case "since":
					return docker.ActionRepositoryTags()
				default:
					return carapace.ActionValues()
				}
			}
		}),
	})

	carapace.Gen(imagesCmd).PositionalCompletion(
		docker.ActionRepositoryTags(),
	)
}
