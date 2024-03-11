package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/flatpak"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:     "search",
	Short:   "Search remote apps/runtimes for text",
	GroupID: "search",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(searchCmd).Standalone()

	searchCmd.Flags().String("arch", "", "Arch to search for")
	searchCmd.Flags().String("columns", "", "What information to show")
	searchCmd.Flags().BoolP("help", "h", false, "Show help options")
	searchCmd.Flags().String("installation", "", "Work on a non-default system-wide installation")
	searchCmd.Flags().Bool("ostree-verbose", false, "Show OSTree debug information")
	searchCmd.Flags().Bool("system", false, "Work on the system-wide installation (default)")
	searchCmd.Flags().BoolP("user", "u", false, "Work on the user installation")
	searchCmd.Flags().BoolP("verbose", "v", false, "Show debug information, -vv for more detail")
	rootCmd.AddCommand(searchCmd)

	carapace.Gen(searchCmd).FlagCompletion(carapace.ActionMap{
		"arch":    flatpak.ActionArches(),
		"columns": columns(flatpak.ActionSearchColumns()),
		// "installation": carapace.ActionValues(), TODO
	})

	carapace.Gen(searchCmd).PositionalCompletion(
		flatpak.ActionApplicationSearch().FilterArgs(),
	)
}
