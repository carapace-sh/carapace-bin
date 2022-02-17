package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/kak_completer/cmd/action"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kak",
	Short: "a vim-inspired, selection oriented code editor",
	Long:  "http://kakoune.org/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	carapace.Override(carapace.Opts{
		LongShorthand: true,
	})
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().String("E", "", "execute argument on server initialisation")
	rootCmd.Flags().String("c", "", "connect to given session")
	rootCmd.Flags().Bool("clear", false, "clear dead sessions")
	rootCmd.Flags().Bool("d", false, "run as a headless session (requires -s)")
	rootCmd.Flags().String("e", "", "execute argument on client initialisation")
	rootCmd.Flags().String("f", "", "filter: for each file, select the entire buffer and execute the given keys")
	rootCmd.Flags().Bool("help", false, "display a help message and quit")
	rootCmd.Flags().String("i", "", "backup the files on which a filter is applied using the given suffix")
	rootCmd.Flags().Bool("l", false, "list existing sessions")
	rootCmd.Flags().Bool("n", false, "do not source kakrc files on startup")
	rootCmd.Flags().String("p", "", "just send stdin as commands to the given session")
	rootCmd.Flags().Bool("q", false, "in filter mode, be quiet about errors applying keys")
	rootCmd.Flags().Bool("ro", false, "readonly mode")
	rootCmd.Flags().String("s", "", "set session name")
	rootCmd.Flags().String("ui", "", "set the type of user interface to use (terminal, dummy, or json)")
	rootCmd.Flags().Bool("version", false, "display kakoune version and exit")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"c":  action.ActionSessions(),
		"p":  action.ActionSessions(),
		"ui": carapace.ActionValues("terminal", "dummy", "json"),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
