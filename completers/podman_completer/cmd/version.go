package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version [options]",
	Short: "Display the Podman version information",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(versionCmd).Standalone()

	versionCmd.Flags().StringP("format", "f", "", "Change the output format to JSON or a Go template")
	rootCmd.AddCommand(versionCmd)
}
