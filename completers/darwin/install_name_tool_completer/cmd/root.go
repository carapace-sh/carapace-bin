package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "install_name_tool",
	Short: "Change dynamic shared library install names",
	Long:  "https://keith.github.io/xcode-manpages/install_name_tool.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("add_rpath", "add_rpath", "", "Add an rpath")
	rootCmd.Flags().StringP("change", "change", "", "Change an install name")
	rootCmd.Flags().StringP("delete_rpath", "delete_rpath", "", "Delete an rpath")
	rootCmd.Flags().BoolP("help", "h", false, "Display usage information")
	rootCmd.Flags().StringP("id", "id", "", "Set the install name")
	rootCmd.Flags().StringP("rpath", "rpath", "", "Change an rpath")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
