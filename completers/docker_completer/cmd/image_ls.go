package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/docker_completer/cmd/action"
	"github.com/spf13/cobra"
)

var image_lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List image",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_lsCmd).Standalone()

	image_lsCmd.Flags().BoolP("all", "a", false, "Show all images (default hides intermediate images)")
	image_lsCmd.Flags().Bool("digests", false, "Show digests")
	image_lsCmd.Flags().StringP("filter", "f", "", "Filter output based on conditions provided")
	image_lsCmd.Flags().String("format", "", "Pretty-print images using a Go template")
	image_lsCmd.Flags().Bool("no-trunc", false, "Don't truncate output")
	image_lsCmd.Flags().BoolP("quiet", "q", false, "Only show numeric IDs")
	imageCmd.AddCommand(image_lsCmd)

	rootAlias(image_lsCmd, func(cmd *cobra.Command, isAlias bool) {
		if isAlias {
			cmd.Use = "images"
		}
		carapace.Gen(cmd).PositionalCompletion(
			action.ActionRepositoryTags(),
		)
	})
}
