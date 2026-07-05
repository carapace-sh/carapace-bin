package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "defaults",
	Short: "access user preferences",
	Long:  "https://keith.github.io/xcode-manpages/defaults.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("currentHost", false, "Operate on the current host's defaults")
	rootCmd.Flags().StringP("host", "host", "", "Operate on a specified host's defaults")
}
