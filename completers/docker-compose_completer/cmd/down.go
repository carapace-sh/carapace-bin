package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var downCmd = &cobra.Command{
	Use:   "down",
	Short: "Stop and remove containers, networks, images, and volumes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	downCmd.Flags().Bool("remove-orphans", false, "Remove containers for services not defined in the")
	downCmd.Flags().String("rmi", "", "Remove images. Type must be one of:")
	downCmd.Flags().StringP("timeout", "t", "", "Specify a shutdown timeout in seconds.")
	downCmd.Flags().BoolP("volumes", "v", false, "Remove named volumes declared in the `volumes`")
	rootCmd.AddCommand(downCmd)

	carapace.Gen(downCmd).FlagCompletion(carapace.ActionMap{
		"rmi": carapace.ActionValuesDescribed(
			"all", "Remove all images used by any service.",
			"local", "Remove only images that don't have a custom tag set by the image field.",
		),
		"volumes": ActionVolumes(),
	})
}
