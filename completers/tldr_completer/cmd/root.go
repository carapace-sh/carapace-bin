package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tldr"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tldr",
	Short: "Python command line client for tldr",
	Long:  "https://github.com/tldr-pages/tldr-python-client",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("color", "c", false, "Override color stripping")
	rootCmd.Flags().BoolP("help", "h", false, "show this help message and exit")
	rootCmd.Flags().StringP("language", "L", "", "Override the default language")
	rootCmd.Flags().BoolP("list", "l", false, "List all available commands for operating system")
	rootCmd.Flags().StringP("platform", "p", "", "Override the operating system")
	rootCmd.Flags().BoolP("render", "r", false, "Render local markdown files")
	rootCmd.Flags().StringP("source", "s", "", "Override the default page source")
	rootCmd.Flags().BoolP("update_cache", "u", false, "Update the local cache of pages and exit")
	rootCmd.Flags().BoolP("version", "v", false, "show program's version number and exit")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"language": os.ActionLanguages(),
		"platform": carapace.ActionValues("linux", "osx", "sunos", "windows", "common"),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		tldr.ActionCommands(),
	)

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin(),
	)
}
