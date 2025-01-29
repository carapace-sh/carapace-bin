package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var container_restoreCmd = &cobra.Command{
	Use:   "restore [options] CONTAINER|IMAGE [CONTAINER|IMAGE...]",
	Short: "Restore one or more containers from a checkpoint",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_restoreCmd).Standalone()

	container_restoreCmd.Flags().BoolP("all", "a", false, "Restore all checkpointed containers")
	container_restoreCmd.Flags().Bool("file-locks", false, "Restore a container with file locks")
	container_restoreCmd.Flags().Bool("ignore-rootfs", false, "Do not apply root file-system changes when importing from exported checkpoint")
	container_restoreCmd.Flags().Bool("ignore-static-ip", false, "Ignore IP address set via --static-ip")
	container_restoreCmd.Flags().Bool("ignore-static-mac", false, "Ignore MAC address set via --mac-address")
	container_restoreCmd.Flags().Bool("ignore-volumes", false, "Do not export volumes associated with container")
	container_restoreCmd.Flags().StringP("import", "i", "", "Restore from exported checkpoint archive (tar.gz)")
	container_restoreCmd.Flags().String("import-previous", "", "Restore from exported pre-checkpoint archive (tar.gz)")
	container_restoreCmd.Flags().BoolP("keep", "k", false, "Keep all temporary checkpoint files")
	container_restoreCmd.Flags().BoolP("latest", "l", false, "Act on the latest container podman is aware of")
	container_restoreCmd.Flags().StringP("name", "n", "", "Specify new name for container restored from exported checkpoint (only works with image or --import)")
	container_restoreCmd.Flags().String("pod", "", "Restore container into existing Pod (only works with image or --import)")
	container_restoreCmd.Flags().Bool("print-stats", false, "Display restore statistics")
	container_restoreCmd.Flags().StringSliceP("publish", "p", []string{}, "Publish a container's port, or a range of ports, to the host (default [])")
	container_restoreCmd.Flags().Bool("tcp-established", false, "Restore a container with established TCP connections")
	containerCmd.AddCommand(container_restoreCmd)
}
