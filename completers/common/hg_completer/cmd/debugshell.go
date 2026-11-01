package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debugshellCmd = &cobra.Command{
	Use:    "debugshell",
	Short:  "run an interactive Python interpreter",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debugshellCmd).Standalone()

	debugshellCmd.Flags().StringP("command", "c", "", "program passed in as a string")
	rootCmd.AddCommand(debugshellCmd)
}
