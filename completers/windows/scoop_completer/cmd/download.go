package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/windows/scoop_completer/cmd/action"
	"github.com/spf13/cobra"
)

var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "download apps in the cache folder and verify hashes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(downloadCmd).Standalone()
	downloadCmd.Flags().StringP("arch", "a", "", "use the specified architecture (32bit|64bit|arm64)")
	downloadCmd.Flags().BoolP("force", "f", false, "force download (overwrite cache)")
	downloadCmd.Flags().BoolP("no-update-scoop", "u", false, "don't update Scoop before downloading")
	downloadCmd.Flags().BoolP("skip-hash-check", "s", false, "skip hash verification (use with caution!)")
	rootCmd.AddCommand(downloadCmd)

	carapace.Gen(downloadCmd).FlagCompletion(carapace.ActionMap{
		"arch": carapace.ActionValues("32bit", "64bit", "arm64"),
	})

	carapace.Gen(downloadCmd).PositionalAnyCompletion(
		action.ActionAvailableApps(),
	)
}
