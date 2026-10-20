package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/mise"
	"github.com/spf13/cobra"
)

var latestCmd = &cobra.Command{
	Use:   "latest",
	Short: "Print the latest version that matches a version request",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(latestCmd).Standalone()

	latestCmd.Flags().String("before", "", "Only consider versions released before a date or at least a duration ago")
	latestCmd.Flags().BoolP("help", "h", false, "Print help")
	latestCmd.Flags().BoolP("installed", "i", false, "Print the latest installed version instead of the latest available one")
	latestCmd.Flags().String("minimum-release-age", "", "Only consider versions released before a date or at least a duration ago")
	latestCmd.Flag("before").Hidden = true
	rootCmd.AddCommand(latestCmd)

	carapace.Gen(latestCmd).PositionalCompletion(
		mise.ActionToolVersions(),
	)
}
