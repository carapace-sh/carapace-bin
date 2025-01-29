package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var container_checkpointCmd = &cobra.Command{
	Use:   "checkpoint [options] CONTAINER [CONTAINER...]",
	Short: "Checkpoint one or more containers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_checkpointCmd).Standalone()

	container_checkpointCmd.Flags().BoolP("all", "a", false, "Checkpoint all running containers")
	container_checkpointCmd.Flags().StringP("compress", "c", "", "Select compression algorithm (gzip, none, zstd) for checkpoint archive.")
	container_checkpointCmd.Flags().String("create-image", "", "Create checkpoint image with specified name")
	container_checkpointCmd.Flags().StringP("export", "e", "", "Export the checkpoint image to a tar.gz")
	container_checkpointCmd.Flags().Bool("file-locks", false, "Checkpoint a container with file locks")
	container_checkpointCmd.Flags().Bool("ignore-rootfs", false, "Do not include root file-system changes when exporting")
	container_checkpointCmd.Flags().Bool("ignore-volumes", false, "Do not export volumes associated with container")
	container_checkpointCmd.Flags().BoolP("keep", "k", false, "Keep all temporary checkpoint files")
	container_checkpointCmd.Flags().BoolP("latest", "l", false, "Act on the latest container podman is aware of")
	container_checkpointCmd.Flags().BoolP("leave-running", "R", false, "Leave the container running after writing checkpoint to disk")
	container_checkpointCmd.Flags().BoolP("pre-checkpoint", "P", false, "Dump container's memory information only, leave the container running")
	container_checkpointCmd.Flags().Bool("print-stats", false, "Display checkpoint statistics")
	container_checkpointCmd.Flags().Bool("tcp-established", false, "Checkpoint a container with established TCP connections")
	container_checkpointCmd.Flags().Bool("with-previous", false, "Checkpoint container with pre-checkpoint images")
	containerCmd.AddCommand(container_checkpointCmd)
}
