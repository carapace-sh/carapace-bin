package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/flatpak"
	"github.com/spf13/cobra"
)

var remoteLsCmd = &cobra.Command{
	Use:     "remote-ls [OPTION…]  [REMOTE or URI]",
	Short:   "Show available runtimes and applications",
	GroupID: "repository",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(remoteLsCmd).Standalone()

	remoteLsCmd.Flags().BoolP("all", "a", false, "List all refs (including locale/debug)")
	remoteLsCmd.Flags().Bool("app", false, "Show only apps")
	remoteLsCmd.Flags().String("app-runtime", "", "List all applications using RUNTIME")
	remoteLsCmd.Flags().String("arch", "", "Limit to this arch (* for all)")
	remoteLsCmd.Flags().Bool("cached", false, "Use local caches even if they are stale")
	remoteLsCmd.Flags().String("columns", "", "What information to show")
	remoteLsCmd.Flags().BoolP("help", "h", false, "Show help options")
	remoteLsCmd.Flags().String("installation", "", "Work on a non-default system-wide installation")
	remoteLsCmd.Flags().Bool("ostree-verbose", false, "Show OSTree debug information")
	remoteLsCmd.Flags().Bool("runtime", false, "Show only runtimes")
	remoteLsCmd.Flags().BoolP("show-details", "d", false, "Show arches and branches")
	remoteLsCmd.Flags().Bool("sideloaded", false, "Only list refs available as sideloads")
	remoteLsCmd.Flags().Bool("system", false, "Work on the system-wide installation (default)")
	remoteLsCmd.Flags().Bool("updates", false, "Show only those where updates are available")
	remoteLsCmd.Flags().BoolP("user", "u", false, "Work on the user installation")
	remoteLsCmd.Flags().BoolP("verbose", "v", false, "Show debug information, -vv for more detail")
	rootCmd.AddCommand(remoteLsCmd)

	// TODO flag
	carapace.Gen(remoteLsCmd).FlagCompletion(carapace.ActionMap{
		// "app-runtime":  carapace.ActionValues(),
		"arch":    flatpak.ActionArches(),
		"columns": columns(flatpak.ActionRemoteContentColumns()),
		// "installation": carapace.ActionValues(),
	})

	// TODO positional
}
