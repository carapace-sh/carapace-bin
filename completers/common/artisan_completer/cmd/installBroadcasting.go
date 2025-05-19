package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var installBroadcastingCmd = &cobra.Command{
	Use:   "install:broadcasting [--composer [COMPOSER]] [--force] [--without-reverb] [--without-node]",
	Short: "Create a broadcasting channel routes file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(installBroadcastingCmd).Standalone()

	installBroadcastingCmd.Flags().String("composer", "", "Absolute path to the Composer binary which should be used to install packages")
	installBroadcastingCmd.Flags().Bool("force", false, "Overwrite any existing broadcasting routes file")
	installBroadcastingCmd.Flags().Bool("without-node", false, "Do not prompt to install Node dependencies")
	installBroadcastingCmd.Flags().Bool("without-reverb", false, "Do not prompt to install Laravel Reverb")
	rootCmd.AddCommand(installBroadcastingCmd)
}
