package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var setInstallLocationCmd = &cobra.Command{
	Use:   "set-install-location",
	Short: "Changes the default install location",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(setInstallLocationCmd).Standalone()

	rootCmd.AddCommand(setInstallLocationCmd)

	carapace.Gen(setInstallLocationCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"0", "auto: let system decide the best location",
			"1", "internal: install on internal device storage",
			"2", "external: install on external media",
		),
	)
}
