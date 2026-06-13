package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pwd",
	Short: "Output the current working directory",
	Long:  "https://fishshell.com/docs/current/cmds/pwd.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("L", "L", false, "logical path")
	rootCmd.Flags().Bool("logical", false, "logical path")
	rootCmd.Flags().BoolS("P", "P", false, "physical path")
	rootCmd.Flags().Bool("physical", false, "physical path")
	rootCmd.Flags().BoolS("h", "h", false, "display help")
}
