package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/flatpak"
	"github.com/spf13/cobra"
)

var remotesCmd = &cobra.Command{
	Use:     "remotes [OPTION…]",
	Short:   "List remote repositories",
	GroupID: "repository",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(remotesCmd).Standalone()

	remotesCmd.Flags().String("columns", "", "What information to show")
	remotesCmd.Flags().BoolP("help", "h", false, "Show help options")
	remotesCmd.Flags().String("installation", "", "Work on a non-default system-wide installation")
	remotesCmd.Flags().Bool("ostree-verbose", false, "Show OSTree debug information")
	remotesCmd.Flags().BoolP("show-details", "d", false, "Show remote details")
	remotesCmd.Flags().Bool("show-disabled", false, "Show disabled remotes")
	remotesCmd.Flags().Bool("system", false, "Work on the system-wide installation (default)")
	remotesCmd.Flags().BoolP("user", "u", false, "Work on the user installation")
	remotesCmd.Flags().BoolP("verbose", "v", false, "Show debug information, -vv for more detail")
	rootCmd.AddCommand(remotesCmd)

	carapace.Gen(remotesCmd).FlagCompletion(carapace.ActionMap{
		"columns":      columns(flatpak.ActionRemoteColumns()),
		"installation": carapace.ActionValues(), // TODO
	})
}
