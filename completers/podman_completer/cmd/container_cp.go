package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var container_cpCmd = &cobra.Command{
	Use:   "cp [options] [CONTAINER:]SRC_PATH [CONTAINER:]DEST_PATH",
	Short: "Copy files/folders between a container and the local filesystem",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_cpCmd).Standalone()

	container_cpCmd.Flags().BoolP("archive", "a", false, "Chown copied files to the primary uid/gid of the destination container.")
	container_cpCmd.Flags().Bool("extract", false, "Deprecated...")
	container_cpCmd.Flags().Bool("overwrite", false, "Allow to overwrite directories with non-directories and vice versa")
	container_cpCmd.Flags().Bool("pause", false, "Deprecated")
	container_cpCmd.Flag("extract").Hidden = true
	container_cpCmd.Flag("pause").Hidden = true
	containerCmd.AddCommand(container_cpCmd)
}
