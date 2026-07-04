package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var whatVersionCmd = &cobra.Command{
	Use:   "what-version",
	Short: "Print the current project version number",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(whatVersionCmd).Standalone()
	whatVersionCmd.Flags().Bool("terse", false, "Print only the version number")
	rootCmd.AddCommand(whatVersionCmd)
}
