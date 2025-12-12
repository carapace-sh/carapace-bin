package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kubeadm",
	Short: "kubeadm: easily bootstrap a secure Kubernetes cluster",
	Long:  "https://kubernetes.io/docs/reference/setup-tools/kubeadm/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.PersistentFlags().String("rootfs", "", "The path to the 'real' host root filesystem. This will cause kubeadm to chroot into the provided path.")
	rootCmd.PersistentFlags().StringP("v", "v", "", "number for the log level verbosity")
	rootCmd.PersistentFlags().String("vmodule", "", "comma-separated list of pattern=N settings for file-filtered logging")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"rootfs": carapace.ActionFiles(),
	})
}
