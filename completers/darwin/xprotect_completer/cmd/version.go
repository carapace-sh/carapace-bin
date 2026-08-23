package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "print the version of the currently installed XProtect assets",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(versionCmd).Standalone()
	rootCmd.AddCommand(versionCmd)

	versionCmd.Flags().Bool("json", false, "Output in JSON format")
}