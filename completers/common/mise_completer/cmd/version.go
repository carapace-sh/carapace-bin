package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:     "version",
	Short:   "Show the version of mise",
	Aliases: []string{"v"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(versionCmd).Standalone()

	versionCmd.Flags().BoolP("help", "h", false, "Print help")
	versionCmd.Flags().BoolP("json", "J", false, "Print the version information in JSON format")
	rootCmd.AddCommand(versionCmd)
}
