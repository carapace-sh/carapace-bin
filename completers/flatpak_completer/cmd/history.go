package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/time"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/flatpak"
	"github.com/spf13/cobra"
)

var historyCmd = &cobra.Command{
	Use:     "history [OPTION…]",
	Short:   "Show history",
	GroupID: "install",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(historyCmd).Standalone()

	historyCmd.Flags().String("columns", "", "What information to show")
	historyCmd.Flags().BoolP("help", "h", false, "Show help options")
	historyCmd.Flags().String("installation", "", "Work on a non-default system-wide installation")
	historyCmd.Flags().Bool("ostree-verbose", false, "Show OSTree debug information")
	historyCmd.Flags().Bool("reverse", false, "Show newest entries first")
	historyCmd.Flags().String("since", "", "Only show changes after TIME")
	historyCmd.Flags().Bool("system", false, "Work on the system-wide installation (default)")
	historyCmd.Flags().String("until", "", "Only show changes before TIME")
	historyCmd.Flags().BoolP("user", "u", false, "Work on the user installation")
	historyCmd.Flags().BoolP("verbose", "v", false, "Show debug information, -vv for more detail")
	rootCmd.AddCommand(historyCmd)

	// TODO flag
	carapace.Gen(historyCmd).FlagCompletion(carapace.ActionMap{
		"columns": columns(flatpak.ActionHistoryColums()),
		// "installation": carapace.ActionValues(),
		"since": time.ActionDateTime(time.DateTimeOpts{}),
		"until": time.ActionDateTime(time.DateTimeOpts{}),
	})
}
