package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var unshareCmd = &cobra.Command{
	Use:   "unshare [options] [COMMAND [ARG...]]",
	Short: "Run a command in a modified user namespace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(unshareCmd).Standalone()

	unshareCmd.Flags().Bool("rootless-netns", false, "Join the rootless network namespace used for CNI and netavark networking")
	rootCmd.AddCommand(unshareCmd)
}
