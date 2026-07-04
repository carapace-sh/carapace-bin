package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/fs"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mkdir",
	Short: "make directories",
	Long:  "https://keith.github.io/xcode-manpages/mkdir.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("mode", "m", "", "Set the file permissions of the final created directory to the specified mode")
	rootCmd.Flags().BoolP("parents", "p", false, "Create any missing intermediate pathname components")
	rootCmd.Flags().BoolP("verbose", "v", false, "Cause mkdir to be verbose, showing files as they are created")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"mode": fs.ActionFileModesNumeric(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionDirectories())
}
