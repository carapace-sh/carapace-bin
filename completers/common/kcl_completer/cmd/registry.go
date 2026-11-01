package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var registryCmd = &cobra.Command{
	Use:   "registry",
	Short: "KCL registry management",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(registryCmd).Standalone()

	registryCmd.PersistentFlags().BoolP("quiet", "q", false, "Set the quiet mode (no output)")
	rootCmd.AddCommand(registryCmd)
}
