package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/tealdeer_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tealdeer",
	Short: "A fast TLDR client",
	Long:  "https://github.com/dbrgn/tealdeer",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute(opts ...func(cmd *cobra.Command)) error {
	for _, opt := range opts {
		opt(rootCmd)
	}
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("clear-cache", "c", false, "Clear the local cache")
	rootCmd.Flags().String("color", "", "Control whether to use color")
	rootCmd.Flags().BoolP("help", "h", false, "Print help information")
	rootCmd.Flags().StringP("language", "L", "", "Override the language")
	rootCmd.Flags().BoolP("list", "l", false, "List all commands in the cache")
	rootCmd.Flags().Bool("no-auto-update", false, "If auto update is configured, disable it for this run")
	rootCmd.Flags().Bool("pager", false, "Use a pager to page output")
	rootCmd.Flags().StringP("platform", "p", "", "Override the operating system")
	rootCmd.Flags().BoolP("quiet", "q", false, "Suppress informational messages")
	rootCmd.Flags().BoolP("raw", "r", false, "Display the raw markdown instead of rendering it")
	rootCmd.Flags().StringP("render", "f", "", "Render a specific markdown file")
	rootCmd.Flags().Bool("seed-config", false, "Create a basic config")
	rootCmd.Flags().Bool("show-paths", false, "Show file and directory paths used by tealdeer")
	rootCmd.Flags().BoolP("update", "u", false, "Update the local cache")
	rootCmd.Flags().BoolP("version", "v", false, "Print the version")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"color":    carapace.ActionValues("always", "auto", "never").StyleF(style.ForKeyword),
		"language": os.ActionLanguages(),
		"platform": carapace.ActionValues("linux", "macos"),
		"render":   carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		action.ActionCommands(),
	)

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin(),
	)
}
