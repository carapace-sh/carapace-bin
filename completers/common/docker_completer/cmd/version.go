package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:     "version [OPTIONS]",
	Short:   "Show the Docker version information",
	GroupID: "common",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(versionCmd).Standalone()

	versionCmd.Flags().StringP("format", "f", "", "Format the output using the given Go template")
	rootCmd.AddCommand(versionCmd)
}
