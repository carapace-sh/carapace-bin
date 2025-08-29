package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/tldr_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
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

	rootCmd.Flags().BoolP("clear-cache", "k", false, "delete the local cache of pages and exit")
	rootCmd.Flags().BoolP("color", "c", false, "override color stripping")
	rootCmd.Flags().StringP("language", "L", "", "override the default language")
	rootCmd.Flags().BoolP("list", "l", false, "list all available commands for operating system")
	rootCmd.Flags().Bool("long-options", false, "display longform options over shortform")
	rootCmd.Flags().BoolP("markdown", "m", false, "just print the plain page file")
	rootCmd.Flags().StringP("platform", "p", "", "override the operating system")
	rootCmd.Flags().Bool("print-completion", false, "print shell completion script")
	rootCmd.Flags().BoolP("render", "r", false, "render local markdown files")
	rootCmd.Flags().String("search", "", "search for a specific command from a query")
	rootCmd.Flags().Bool("short-options", false, "display shortform options over longform")
	rootCmd.Flags().StringP("source", "s", "", "override the default page source")
	rootCmd.Flags().BoolP("update", "u", false, "update the local cache of pages and exit")
	rootCmd.Flags().Bool("update_cache", false, "update the local cache of pages and exit")
	rootCmd.Flags().BoolP("version", "v", false, "show program's version number and exit")

	rootCmd.MarkFlagsMutuallyExclusive("update", "update_cache")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"language": os.ActionLanguages(),
		"platform": carapace.ActionValues("android", "freebsd", "linux", "netbsd", "openbsd", "osx", "sunos", "windows", "common"),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionCommands(rootCmd.Flag("language").Value.String())
		}),
	)

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin(),
	)
}
