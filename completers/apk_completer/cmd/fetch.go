package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/apk_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/completers/apk_completer/cmd/common"
	"github.com/spf13/cobra"
)

var fetchCmd = &cobra.Command{
	Use:     "fetch",
	Short:   "Download packages from repositories to a local directory",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: "repository maintenance",
}

func init() {
	carapace.Gen(fetchCmd).Standalone()

	fetchCmd.Flags().String("built-after", "", "Only fetch packages that have buildtime more recent than TIMESPEC")
	fetchCmd.Flags().BoolP("link", "l", false, "Create hard links if possible")
	fetchCmd.Flags().StringP("output", "o", "", "Write the downloaded file(s) to DIR")
	fetchCmd.Flags().BoolP("recursive", "R", false, "Fetch packages and all of their dependencies")
	fetchCmd.Flags().Bool("simulate", false, "Simulate the requested operation without making any changes")
	fetchCmd.Flags().BoolP("stdout", "s", false, "Dump the .apk file(s) to stdout")
	fetchCmd.Flags().Bool("url", false, "Print the full URL for downloaded packages")
	fetchCmd.Flags().BoolP("world", "w", false, "Download packages needed to satisfy WORLD")
	common.AddGlobalFlags(fetchCmd)
	common.AddSourceFlags(fetchCmd)
	rootCmd.AddCommand(fetchCmd)

	carapace.Gen(fetchCmd).FlagCompletion(carapace.ActionMap{
		"output": carapace.ActionDirectories(),
	})

	carapace.Gen(fetchCmd).PositionalAnyCompletion(
		action.ActionPackageSearch(fetchCmd).FilterArgs(),
	)
}
