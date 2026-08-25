package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pi"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update pi, extensions, or model catalogs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(updateCmd).Standalone()
	updateCmd.Flags().Bool("all", false, "Update pi and installed packages")
	updateCmd.Flags().BoolP("approve", "a", false, "Trust project-local files for this command")
	updateCmd.Flags().String("extension", "", "Update one package only")
	updateCmd.Flags().Bool("extensions", false, "Update installed packages only")
	updateCmd.Flags().Bool("force", false, "Reinstall pi even if the current version is latest")
	updateCmd.Flags().Bool("models", false, "Refresh model catalogs only")
	updateCmd.Flags().Bool("no-approve", false, "Ignore project-local files for this command")
	updateCmd.Flags().Bool("self", false, "Update pi only")
	rootCmd.AddCommand(updateCmd)

	carapace.Gen(updateCmd).PositionalCompletion(
		carapace.Batch(
			carapace.ActionValues("self", "pi"),
			pi.ActionPackages(),
		).ToA(),
	)
}