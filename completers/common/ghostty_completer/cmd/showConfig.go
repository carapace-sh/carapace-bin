package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var showConfigCmd = &cobra.Command{
	Use:   "+show-config",
	Short: "show the current configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(showConfigCmd).Standalone()

	showConfigCmd.Flags().Bool("changes-only", false, "Only show the options that have been changed")
	showConfigCmd.Flags().Bool("default", false, "Show the default configuration")
	showConfigCmd.Flags().Bool("docs", false, "Print the documentation above each option as a comment")
	showConfigCmd.Flags().Bool("help", false, "show help")
	rootCmd.AddCommand(showConfigCmd)
}
