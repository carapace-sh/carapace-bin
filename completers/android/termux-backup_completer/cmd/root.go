package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "termux-backup",
	Short: "Backup the Termux prefix to a tar archive",
	Long:  "https://github.com/termux/termux-tools/blob/master/scripts/termux-backup.in",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("force", "f", false, "force overwrite of existing output file without prompting")
	rootCmd.Flags().BoolP("help", "h", false, "show usage and exit")
	rootCmd.Flags().Bool("ignore-read-failure", false, "suppress read permission denials")
	rootCmd.Flags().BoolP("usage", "?", false, "show usage and exit")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
