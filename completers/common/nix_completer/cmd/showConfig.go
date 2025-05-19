package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var showConfigCmd = &cobra.Command{
	Use:     "show-config",
	Short:   "show the Nix configuration or the value of a specific setting",
	GroupID: "utility",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(showConfigCmd).Standalone()

	showConfigCmd.Flags().Bool("json", false, "Produce output in JSON format")
	rootCmd.AddCommand(showConfigCmd)

	addLoggingFlags(showConfigCmd)

	// TODO positional completion
}
