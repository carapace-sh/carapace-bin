package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var shellCmd = &cobra.Command{
	Use:   "shell",
	Short: "Start a new shell or run a command with access to your packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(shellCmd).Standalone()
	shellCmd.Flags().StringP("config", "c", "", "path to directory containing a devbox.json config file")
	shellCmd.Flags().Bool("print-env", false, "Print script to setup shell environment")
	rootCmd.AddCommand(shellCmd)

	carapace.Gen(shellCmd).FlagCompletion(carapace.ActionMap{
		"config": carapace.ActionDirectories(),
	})

	carapace.Gen(shellCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)

	// TODO dash completion - if possible?
}
