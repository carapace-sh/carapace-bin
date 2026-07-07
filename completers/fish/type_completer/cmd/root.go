package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "type",
	Short: "Locate a command",
	Long:  "https://fishshell.com/docs/current/cmds/type.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("P", "P", false, "PATH only")
	rootCmd.Flags().BoolS("a", "a", false, "all definitions")
	rootCmd.Flags().Bool("all", false, "all definitions")
	rootCmd.Flags().String("color", "", "color output")
	rootCmd.Flags().BoolS("f", "f", false, "no functions")
	rootCmd.Flags().Bool("force-path", false, "PATH only")
	rootCmd.Flags().BoolS("h", "h", false, "display help")
	rootCmd.Flags().Bool("no-functions", false, "no functions")
	rootCmd.Flags().BoolS("p", "p", false, "print path")
	rootCmd.Flags().Bool("path", false, "print path")
	rootCmd.Flags().BoolS("q", "q", false, "suppress output")
	rootCmd.Flags().Bool("query", false, "suppress output")
	rootCmd.Flags().Bool("quiet", false, "suppress output")
	rootCmd.Flags().BoolS("s", "s", false, "short output")
	rootCmd.Flags().Bool("short", false, "short output")
	rootCmd.Flags().BoolS("t", "t", false, "print type")
	rootCmd.Flags().Bool("type", false, "print type")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"color": carapace.ActionValues("always", "never", "auto"),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionExecutables().FilterArgs(),
	)
}
