package cmd

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/glow"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "glow [SOURCE|DIR]",
	Short: "Render markdown on the CLI, with pizzazz!",
	Long:  "https://github.com/charmbracelet/glow",
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

	rootCmd.Flags().BoolP("all", "a", false, "show system files and directories (TUI-mode only)")
	rootCmd.PersistentFlags().String("config", "", "config file")
	rootCmd.Flags().BoolP("line-numbers", "l", false, "show line numbers (TUI-mode only)")
	rootCmd.Flags().BoolP("mouse", "m", false, "enable mouse wheel (TUI-mode only)")
	rootCmd.Flags().BoolP("pager", "p", false, "display with pager")
	rootCmd.Flags().BoolP("preserve-new-lines", "n", false, "preserve newlines in the output")
	rootCmd.Flags().StringP("style", "s", "", "style name or JSON path")
	rootCmd.Flags().StringP("width", "w", "", "word-wrap at width (set to 0 to disable)")
	rootCmd.Flag("mouse").Hidden = true

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"config": carapace.ActionFiles(".yml", ".yaml"),
		"style": carapace.Batch(
			glow.ActionStyles(),
			carapace.ActionFiles(".json"),
		).ToA(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.Batch(
			carapace.ActionCallback(func(c carapace.Context) carapace.Action {
				if !strings.HasPrefix(c.Value, "https://") {
					return carapace.ActionFiles(".md")
				}
				return carapace.ActionValues()
			}),
			git.ActionRepositorySearch(git.SearchOpts{}.Default()),
		).ToA(),
	)
}
