package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/tlp_completer/cmd/action"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tlp",
	Short: "apply laptop power saving settings",
	Long:  "https://linrunner.de/tlp",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("version", false, "Print TLP version")

	carapace.Gen(rootCmd).DashAnyCompletion(
		action.ActionConfigParams(),
	)
}
