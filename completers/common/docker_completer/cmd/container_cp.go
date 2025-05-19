package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
	"github.com/carapace-sh/carapace/pkg/condition"
	"github.com/spf13/cobra"
)

var container_cpCmd = &cobra.Command{
	Use:   "cp [OPTIONS] CONTAINER:SRC_PATH DEST_PATH|-",
	Short: "Copy files/folders between a container and the local filesystem",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_cpCmd).Standalone()

	container_cpCmd.Flags().BoolP("archive", "a", false, "Archive mode (copy all uid/gid information)")
	container_cpCmd.Flags().BoolP("follow-link", "L", false, "Always follow symbol link in SRC_PATH")
	container_cpCmd.Flags().BoolP("quiet", "q", false, "Suppress progress output during copy. Progress output is automatically suppressed if no terminal is attached")
	containerCmd.AddCommand(container_cpCmd)

	carapace.Gen(container_cpCmd).PositionalCompletion(
		carapace.Batch(
			carapace.ActionFiles(),
			docker.ActionContainerPath().UnlessF(condition.CompletingPath),
		).ToA(),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if condition.File(c.Args[0])(c) {
				return docker.ActionContainerPath()
			}
			return carapace.ActionFiles()
		}),
	)
}
