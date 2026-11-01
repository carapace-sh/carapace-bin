package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var downgradeCmd = &cobra.Command{
	Use:   "downgrade",
	Short: "Downgrade Flutter to the last active version for the current channel",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(downgradeCmd).Standalone()

	downgradeCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	downgradeCmd.Flags().Bool("no-prompt", false, "Show the downgrade prompt.")
	downgradeCmd.Flags().Bool("prompt", false, "Show the downgrade prompt.")
	downgradeCmd.Flags().String("working-directory", "", "Override the downgrade working directory. This is only intended to enable integration testing of the tool itself. It allows one to use the flutter tool from one checkout to downgrade a different checkout.")
	rootCmd.AddCommand(downgradeCmd)
	carapace.Gen(downgradeCmd).FlagCompletion(carapace.ActionMap{
		"working-directory": carapace.ActionDirectories(),
	})
}
