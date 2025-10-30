package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var projectCmd = &cobra.Command{
	Use:     "project",
	Short:   "List and manipulate projects",
	Aliases: []string{"projects"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(projectCmd).Standalone()

	projectCmd.Flags().StringP("app-data-dir", "d", "", "The location of the directory to contain app data")
	projectCmd.Flags().StringP("app-suffix", "s", "", "A suffix like `dev` to refer to projects of the development version of the application")
	projectCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(projectCmd)
}
