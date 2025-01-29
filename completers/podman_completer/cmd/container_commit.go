package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var container_commitCmd = &cobra.Command{
	Use:   "commit [options] CONTAINER [IMAGE]",
	Short: "Create new image based on the changed container",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_commitCmd).Standalone()

	container_commitCmd.Flags().StringP("author", "a", "", "Set the author for the image committed")
	container_commitCmd.Flags().StringSliceP("change", "c", []string{}, "Apply the following possible instructions to the created image (default []): CMD | ENTRYPOINT | ENV | EXPOSE | LABEL | ONBUILD | STOPSIGNAL | USER | VOLUME | WORKDIR")
	container_commitCmd.Flags().String("config", "", "`file` containing a container configuration to merge into the image")
	container_commitCmd.Flags().StringP("format", "f", "", "`Format` of the image manifest and metadata")
	container_commitCmd.Flags().String("iidfile", "", "`file` to write the image ID to")
	container_commitCmd.Flags().Bool("include-volumes", false, "Include container volumes as image volumes")
	container_commitCmd.Flags().StringP("message", "m", "", "Set commit message for imported image")
	container_commitCmd.Flags().BoolP("pause", "p", false, "Pause container during commit")
	container_commitCmd.Flags().BoolP("quiet", "q", false, "Suppress output")
	container_commitCmd.Flags().BoolP("squash", "s", false, "squash newly built layers into a single new layer")
	containerCmd.AddCommand(container_commitCmd)
}
