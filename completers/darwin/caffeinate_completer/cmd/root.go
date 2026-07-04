package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "caffeinate",
	Short: "prevent the system from sleeping",
	Long:  "https://keith.github.io/xcode-manpages/caffeinate.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("display", "d", false, "Create an assertion to prevent the display from sleeping")
	rootCmd.Flags().StringP("pid", "w", "", "Wait for the process with the specified pid to exit")
	rootCmd.Flags().BoolP("prevent-idle", "i", false, "Create an assertion to prevent the system from idle sleeping")
	rootCmd.Flags().BoolP("prevent-sleep", "s", false, "Create an assertion to prevent the system from sleeping")
	rootCmd.Flags().BoolP("prevent-user", "u", false, "Create an assertion to declare that user is active")
	rootCmd.Flags().StringP("timeout", "t", "", "Specifies the timeout value in seconds")

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
