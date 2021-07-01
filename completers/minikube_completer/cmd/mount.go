package cmd

import (
	"github.com/spf13/cobra"
)

var mountCmd = &cobra.Command{
	Use:   "mount",
	Short: "Mounts the specified directory into minikube",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	mountCmd.Flags().String("9p-version", "9p2000.L", "Specify the 9p version that the mount should use")
	mountCmd.Flags().String("gid", "docker", "Default group id used for the mount")
	mountCmd.Flags().String("ip", "", "Specify the ip that the mount should be setup on")
	mountCmd.Flags().Bool("kill", false, "Kill the mount process spawned by minikube start")
	mountCmd.Flags().Uint("mode", 493, "File permissions used for the mount")
	mountCmd.Flags().Int("msize", 262144, "The number of bytes to use for 9p packet payload")
	mountCmd.Flags().StringSlice("options", []string{}, "Additional mount options, such as cache=fscache")
	mountCmd.Flags().String("type", "9p", "Specify the mount filesystem type (supported types: 9p)")
	mountCmd.Flags().String("uid", "docker", "Default user id used for the mount")
	rootCmd.AddCommand(mountCmd)
}
