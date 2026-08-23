package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "wdutil",
	Short: "Wireless Diagnostics command line utility",
	Long:  "https://keith.github.io/xcode-manpages/wdutil.8.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.AddCommand(cleanCmd)
	rootCmd.AddCommand(diagnoseCmd)
	rootCmd.AddCommand(dumpCmd)
	rootCmd.AddCommand(infoCmd)
	rootCmd.AddCommand(logCmd)
	rootCmd.AddCommand(privateMACCmd)
}
