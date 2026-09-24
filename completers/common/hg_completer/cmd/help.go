package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/hg"
	"github.com/spf13/cobra"
)

var helpCmd = &cobra.Command{
	Use:     "help",
	Short:   "show help for a given topic or a help overview",
	GroupID: groups[group_help].ID,
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(helpCmd).Standalone()

	helpCmd.Flags().BoolP("command", "c", false, "show only help for commands")
	helpCmd.Flags().BoolP("extension", "e", false, "show only help for extensions")
	helpCmd.Flags().BoolP("keyword", "k", false, "show topics matching keyword")
	helpCmd.Flags().StringArrayP("system", "s", nil, "show help for specific platform(s)")
	rootCmd.AddCommand(helpCmd)

	carapace.Gen(helpCmd).FlagCompletion(carapace.ActionMap{
		"system": carapace.ActionValues("aix", "cygwin", "darwin", "dragonflybsd", "freebsd", "irix", "linux", "netbsd", "openbsd", "plan9", "posix", "unix", "vms", "windows", "win32", "verbose"),
	})

	carapace.Gen(helpCmd).PositionalAnyCompletion(
		hg.ActionTopics(),
	)
}
